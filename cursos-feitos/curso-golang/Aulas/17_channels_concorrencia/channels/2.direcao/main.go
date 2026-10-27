package main

import (
	"fmt"
	"time"
)

func enviar(ch chan<- int) {
	fmt.Println("Enviando valor para o channel")
	ch <- 100 // set
	fmt.Println("valor enviado com sucesso")
}

func receber(ch <-chan int) {
	time.Sleep(time.Second)
	valor := <-ch // get -> leitura
	fmt.Println("valor recebido", valor)
}

func main() {
	// comunicação entre goroutines

	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go enviar(ch)
	}
	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		go receber(ch)
	}
	time.Sleep(time.Second * 5)
}
