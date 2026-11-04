package main

import "fmt"

func produzir(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go produzir(ch)

	for v := range ch {
		fmt.Println("recebi", v)
	}
	fmt.Println("canal fechado, laço terminou")

	v, ok := <-ch
	fmt.Println("de canal fechado:", v, ok)
}
