package main

import "fmt"

func produtor(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go produtor(ch)

	soma := 0
	for valor := range ch {
		soma += valor
		fmt.Printf("dentro do laço: %d\n", valor)
	}

	fmt.Println("depois do laço:", soma)
}
