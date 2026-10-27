package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func testeMutex(iteracoes int) (int64, time.Duration) {
	var (
		contador int64
		mu       sync.Mutex
		wg       sync.WaitGroup
		start    = time.Now()
	)

	for i := 0; i < iteracoes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			contador++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return contador, time.Since(start)
}

func testeAtomic(iteracoes int) (int64, time.Duration) {
	var (
		contador int64
		wg       sync.WaitGroup
		start    = time.Now()
	)
	for i := 0; i < iteracoes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&contador, 1)
		}()
	}
	wg.Wait()
	return contador, time.Since(start)
}

func porcentagem(t1, t2 time.Duration) float64 {
	return (float64(t1-t2) / float64(t1)) * 100.0
}

func main() {
	iteracoes := 1_000_000

	fmt.Printf("Executando com %d iterações...\n", iteracoes)

	c1, t1 := testeMutex(iteracoes)
	fmt.Printf("[Mutext] Contador final: %d, Tempo: %s\n", c1, t1)

	c2, t2 := testeAtomic(iteracoes)
	fmt.Printf("[Atomic] Contador final: %d, Tempo: %s\n", c2, t2)

	fmt.Printf("[Resultado] %.2f%% mais rápido\n", porcentagem(t1, t2))
}
