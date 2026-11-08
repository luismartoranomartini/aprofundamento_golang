package main

import "fmt"

func produzir(out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- i
	}
	close(out)
}

func consumir(in <-chan int) {
	for v := range in {
		fmt.Println("recebi", v)
	}
}

func main() {
	ch := make(chan int)
	go produzir(ch)
	consumir(ch)
}
