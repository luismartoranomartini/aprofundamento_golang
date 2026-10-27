package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada!")

	for msg := range canal {
		fmt.Println(msg)
	}

}

func escrever(txt string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- txt
		time.Sleep(time.Second)
	}
	close(canal)

}
