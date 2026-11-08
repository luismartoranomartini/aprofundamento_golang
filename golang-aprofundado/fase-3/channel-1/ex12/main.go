package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println("vou receber de um canal nil...")
	<-ch
}
