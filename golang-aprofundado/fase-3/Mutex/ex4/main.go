package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	config string
	cargas atomic.Int64
	uma    sync.Once
)

func obterConfig() string {
	uma.Do(func() {
		cargas.Add(1)
		config = "config carregada"
	})
	return config
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = obterConfig()
		}()
	}
	wg.Wait()
	fmt.Println("valor:", config)
	fmt.Println("vezes que carregou:", cargas.Load())
}
