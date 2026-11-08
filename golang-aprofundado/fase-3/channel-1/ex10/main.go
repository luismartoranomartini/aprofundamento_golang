package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	fmt.Println("vou fechar de novo")
	close(ch)
}
