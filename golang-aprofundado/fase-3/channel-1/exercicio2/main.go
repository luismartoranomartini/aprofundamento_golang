package main

import "fmt"

func dobrar(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i + i
	}

	close(out)
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := 1; i <= 4; i++ {
			in <- i
		}
		close(in)
	}()
	go dobrar(in, out)

	for v := range out {
		fmt.Println(v)
	}
}
