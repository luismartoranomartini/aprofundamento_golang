package main

import "fmt"

func dizer(msg string) {
	fmt.Println(msg)
}

func main() {
	go dizer("Olá da gouroutine")
	fmt.Println("fim do main")
}
