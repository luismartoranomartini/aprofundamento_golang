package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var cargas atomic.Int64
	carregar := sync.OnceValue(func() string {
		cargas.Add(1)
		return "config carregada"
	})
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = carregar()
		}()
	}
	wg.Wait()
	fmt.Println("valor:", carregar())
	fmt.Println("vezes que carregou", cargas.Load())
}
