package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	close(ch)
	fmt.Println("vou enviar em canal fechado")
	ch <- 1
}
