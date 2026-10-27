package main

import (
	"fmt"
	"sync"
)

// var mu sync.Mutex

func main() {
	var (
		m  sync.Map
		wg sync.WaitGroup
	)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			m.Store(i, i)
		}()
	}
	wg.Wait()

	m.Range(func(chave, valor any) bool {
		fmt.Println(chave, valor)
		return true
	})

	valor, ok := m.Load(50)
	if ok {
		fmt.Println("Valor encontado", valor)
	} else {
		fmt.Println("valor nãp encotrado")
	}
	teste(m)
}

func teste(m map[any]any) {

}
