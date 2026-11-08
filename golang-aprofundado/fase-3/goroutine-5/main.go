package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs lógicas:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:  ", runtime.GOMAXPROCS(0)) // 0 = só consulta

	// Foto da memória de pilha ANTES de criar as goroutines
	var antes, depois runtime.MemStats
	runtime.ReadMemStats(&antes)

	const n = 100_000
	parar := make(chan struct{}) // canal usado só para segurar as goroutines
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-parar // fica bloqueada, esperando o canal fechar
		}()
	}
	fmt.Println("goroutines vivas:", runtime.NumGoroutine())

	// Foto DEPOIS: a diferença dividida por n é a pilha de cada goroutine
	runtime.ReadMemStats(&depois)
	fmt.Printf("bytes de pilha por goroutine: %d\n", (depois.StackInuse-antes.StackInuse)/n)

	close(parar) // libera todas de uma vez
	wg.Wait()
}
