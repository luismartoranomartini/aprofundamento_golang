package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type itemTarefa struct {
	indice int
	valor  int
}

type itemResultado struct {
	indice int
	valor  int
}

// Processar executa n workers para processar tarefas concorrentemente.
// Mantém a ordem original, respeita o contexto e impede vazamento de goroutines.
func Processar(ctx context.Context, tarefas []int, n int) ([]int, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	total := len(tarefas)
	if total == 0 {
		return []int{}, nil
	}

	if n > total {
		n = total
	}
	if n <= 0 {
		n = 1
	}

	canalTarefas := make(chan itemTarefa)
	canalResultados := make(chan itemResultado)

	var wg sync.WaitGroup

	// 1. Inicia n workers
	for w := 0; w < n; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case t, ok := <-canalTarefas:
					if !ok {
						return
					}

					// Simula carga de trabalho
					select {
					case <-time.After(20 * time.Millisecond):
					case <-ctx.Done():
						return
					}

					// Envio do resultado com cancelamento cooperativo
					select {
					case <-ctx.Done():
						return
					case canalResultados <- itemResultado{indice: t.indice, valor: t.valor * t.valor}:
					}
				}
			}
		}()
	}

	// 2. Goroutine produtora: distribui os itens e fecha canalTarefas
	go func() {
		defer close(canalTarefas)
		for i, v := range tarefas {
			select {
			case <-ctx.Done():
				return
			case canalTarefas <- itemTarefa{indice: i, valor: v}:
			}
		}
	}()

	// 3. Goroutine coordenadora: fecha canalResultados quando todos os workers terminarem
	go func() {
		wg.Wait()
		close(canalResultados)
	}()

	// 4. Consumo, preservação da ordem e validação de término
	resultados := make([]int, total)
	recebidos := 0

	for {
		select {
		case <-ctx.Done():
			wg.Wait() // Assegura que nenhum worker fique ativo antes de sair
			return nil, ctx.Err()

		case res, ok := <-canalResultados:
			if !ok {
				// O canal fechou. Se o select caiu aqui por sorteio pseudoaleatório
				// ou cancelamento súbito antes de coletar tudo, NÃO retorne nil.
				if recebidos < total {
					return nil, ctx.Err()
				}
				return resultados, nil
			}

			resultados[res.indice] = res.valor
			recebidos++
			if recebidos == total {
				return resultados, nil
			}
		}
	}
}

func main() {
	// Cenário 1: Sucesso
	ctx := context.Background()
	tarefas := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res, err := Processar(ctx, tarefas, 3)
	fmt.Printf("Cenário normal -> Res: %v | Err: %v\n", res, err)

	// Cenário 2: Cancelamento via Timeout (sem retorno de slice incompleto com err == nil)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()

	resCancelado, errCancelado := Processar(ctxTimeout, tarefas, 3)
	fmt.Printf("Cenário timeout -> Res: %v | Err: %v\n", resCancelado, errCancelado)
}
