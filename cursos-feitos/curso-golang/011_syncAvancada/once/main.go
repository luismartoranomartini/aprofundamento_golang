package main

import (
	"fmt"
	"sync"
	"time"
)

type BancoDeDados struct {
	conectado bool
}

var (
	bancoDeDados *BancoDeDados
	once         sync.Once
)

func ObterBancoDeDados() *BancoDeDados {
	once.Do(func() {
		fmt.Println("Conectando ao Banco de Dados")
		time.Sleep(time.Second * 2)
		bancoDeDados = &BancoDeDados{conectado: true}
		fmt.Println("Conectado!")
	})
	return bancoDeDados
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("Goroutine %d tentando acessar o banco...\n", i)
			banco := ObterBancoDeDados()
			fmt.Printf("Goroutine %d acessou o banco. Conectado=%v\n", i, banco.conectado)
		}()
	}
	wg.Wait()
	banco := ObterBancoDeDados()
	fmt.Printf("Main acessou o banco. Conectado=%v\n", banco.conectado)
	fmt.Println("Conectado!")
	fmt.Println("Fim")

}
