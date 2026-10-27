package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	var (
		m  = make(map[int]int)
		mu sync.Mutex
		wg sync.WaitGroup
	)
	fmt.Println(m)
	fmt.Println(mu)
	fmt.Println(wg)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mu.Lock() // trava
			m[i] = i
			mu.Unlock() // destrava

			//wg.Done() // decremento
		}()
	}
	// wg.Wait()
	// for chave, valor := range m {
	// 	fmt.Println(chave, valor)
	// }

	fmt.Println("Fim da execução")
}
