package main

import (
	"errors"
	"fmt"
	"sync"
)

type Conta struct {
	mu    sync.RWMutex
	saldo int
}

func (c *Conta) Depositar(v int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.saldo += v
}

func (c *Conta) Sacar(v int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v > c.saldo {
		return errors.New("Saque não pode ser realizado")
	}
	c.saldo -= v
	return nil
}

func (c *Conta) Saldo() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.saldo
}

func main() {
	// c := &Conta{}
	// var wg sync.WaitGroup
	//
	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		c.Depositar(10)
	// 	}()
	// }
	//
	// wg.Wait()
	//
	// var erros atomic.Int64
	//
	// for i := 0; i < 150; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		if err := c.Sacar(10); err != nil {
	// 			erros.Add(1)
	// 		}
	// 	}()
	// }
	// wg.Wait()
	// fmt.Printf("Saldo final: %d\n", c.Saldo())
	//
	// fmt.Printf("Erro de saque: %d\n", erros.Load())

	var wg sync.WaitGroup
	contador := 0 // variável comum, sem proteção atômica ou mutex

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador++ // Linha 17: leitura e escrita concorrentes
		}()
	}

	wg.Wait()
	fmt.Println("Contador final:", contador)
}
