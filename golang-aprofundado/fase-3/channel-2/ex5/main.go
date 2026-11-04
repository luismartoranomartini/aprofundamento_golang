package main

import (
	"fmt"
	"time"
)

func trabalhador(parar <-chan struct{}) {
	n := 0
	for {
		select {
		case <-parar:
			fmt.Println("tratrabalhador: parando")
			return
		case <-time.After(100 * time.Millisecond):
			n++
			fmt.Println("trabalhador: tique", n)
		}
	}
}

func main() {
	parar := make(chan struct{})
	fim := make(chan struct{})

	go func() {
		trabalhador(parar)
		close(fim)
	}()

	time.Sleep(350 * time.Millisecond)
	close(parar)
	<-fim
	fmt.Println("main: fim")
}
