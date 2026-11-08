package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	saldo := 100
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			saldo -= 10
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(saldo)
}
