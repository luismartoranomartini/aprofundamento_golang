package main

import "fmt"

func main() {
	canal := make(chan string, 3)

	canal <- "Olá, mundo!"
	canal <- "Programando em Go"
	canal <- "Programando em Golang"

	msg := <-canal
	msg2 := <-canal
	fmt.Println(msg)
	fmt.Println(msg2)

}
