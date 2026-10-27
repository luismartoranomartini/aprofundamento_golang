package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

func semPool(iteracoes int) time.Duration {
	start := time.Now()

	for i := 0; i < iteracoes; i++ {
		buffer := bytes.Buffer{}
		buffer.WriteString("Olá, mundo!")
		_ = buffer.String()
	}
	return time.Since(start)
}

func comPool(iteracoes int) time.Duration {
	start := time.Now()
	pool := sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	for i := 0; i < iteracoes; i++ {
		buffer := pool.Get().(*bytes.Buffer)

		buffer.WriteString("Olá, mundo!")
		_ = buffer.String()

		buffer.Reset()
		pool.Put(buffer)
	}
	return time.Since(start)
}

func porcentagem(t1, t2 time.Duration) float64 {
	return (float64(t1-t2) / float64(t1)) * 100.0
}

func main() {
	iteracoes := 100_000_000

	fmt.Printf("Executando comparação com %d iteraçoes...\n", iteracoes)

	t1 := semPool(iteracoes)
	fmt.Printf("[Sem Pool] Tempo total: %s\n", t1)

	t2 := comPool(iteracoes)
	fmt.Printf("[Com Pool] Tempo total: %s\n", t2)

	fmt.Printf("[Resultado] %.2f%% mais rápido\n", porcentagem(t1, t2))

}
