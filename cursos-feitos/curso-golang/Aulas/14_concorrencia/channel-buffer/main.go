package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int, 100)
	go setList(channel)

	// valor := <- channel //Lendo Enviando

	for v := range channel {
		fmt.Println("Recebendo: ", v)
		time.Sleep(time.Second)
	}

}

func setList(channel chan <- int) {
	for i := 0; i < 10; i++ {
		channel <- i //Escrevendo Recebendo
		fmt.Println("Enviando: ", i)
	}
	close(channel)
}