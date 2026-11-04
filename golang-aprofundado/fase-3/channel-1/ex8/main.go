package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	resultados := make(chan int)

	var wg sync.WaitGroup
	for _, n := range numeros {
		wg.Add(1)
		go func() {
			time.Sleep(500 * time.Millisecond)
			resultados <- n * n
		}()
	}

	go func() {
		wg.Wait()
		close(resultados)
	}()

	for range numeros {
		fmt.Println(<-resultados)
	}
}
