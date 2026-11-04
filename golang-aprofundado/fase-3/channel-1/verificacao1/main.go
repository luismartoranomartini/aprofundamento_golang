package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)
	fmt.Println(<-ch, <-ch, <-ch)
	v, ok := <-ch
	fmt.Println(v, ok)
}
