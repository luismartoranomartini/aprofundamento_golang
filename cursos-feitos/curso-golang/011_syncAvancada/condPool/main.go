package main

import (
	"fmt"
	"sync"
	"time"
)

type PoolDeConexao struct {
	conexao []int
	mu      sync.Mutex
	cond    *sync.Cond
}

func NovoPoolDeConexoes(tamanho int) *PoolDeConexao {
	pool := &PoolDeConexao{
		conexao: make([]int, tamanho),
	}
	// for i := 0; i < tamanho; i++ {
	// 	pool.conexao[i] = i + 1
	// }
	for i := range tamanho {
		pool.conexao[i] = i + 1
	}
	pool.cond = sync.NewCond(&pool.mu)
	return pool
}

func (p *PoolDeConexao) ObterConeexao() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conexao) == 0 {
		fmt.Println("Nenhuma conexao disponível. Aguardando...")
		p.cond.Wait()
	}
	conexao := p.conexao[0]
	p.conexao = p.conexao[1:]
	fmt.Printf("Conexao %d adquirida.\n", conexao)
	return conexao
}

func (p *PoolDeConexao) LiberarConexao(conexao int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.conexao = append(p.conexao, conexao)
	fmt.Printf("Conexao %d liberada\n", conexao)

	p.cond.Signal()
}

func main() {
	PoolDeConexao := NovoPoolDeConexoes(2)

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 1; i <= 5; i++ {

		go func() {
			defer wg.Done()

			conexao := PoolDeConexao.ObterConeexao()
			fmt.Printf("Goroutine %d utilizando a conexao %d\n", i, conexao)
			time.Sleep(time.Second * 2)
			PoolDeConexao.LiberarConexao(conexao)
		}()
	}
	wg.Wait()
	fmt.Println("Fim")
}
