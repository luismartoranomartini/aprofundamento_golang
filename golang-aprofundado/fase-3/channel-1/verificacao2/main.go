package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	v := <-ch
	fmt.Println(v)
	ch <- 2
	fmt.Println("fim")
}
