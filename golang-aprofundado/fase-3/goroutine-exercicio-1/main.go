package main

import (
	"fmt"
	"sync"
	"time"
)

type Prato struct {
	nome  string
	tempo time.Duration
}

func preparar(prato string, d time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(d)
	fmt.Printf("%s pronto\n", prato)
}

func main() {
	cardapio := []Prato{
		{"sopa", 1 * time.Second},
		{"massa", 2 * time.Second},
		{"sobrenome", 3 * time.Second},
	}
	var wg sync.WaitGroup
	inicio := time.Now()

	for _, p := range cardapio {
		wg.Add(1)
		go preparar(p.nome, p.tempo, &wg)
	}
	wg.Wait()
	fmt.Printf("tudo pronto em %v\n", time.Since(inicio).Round(time.Second))
}
