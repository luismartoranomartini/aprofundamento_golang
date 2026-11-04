package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Println("len:", len(ch), "cap:", cap(ch))

	ch <- 10
	ch <- 20
	fmt.Println("len:", len(ch), "cap:", cap(ch))

	fmt.Println(<-ch)
	fmt.Println("len:", len(ch), "cap:", cap(ch))
	fmt.Println(<-ch)
	fmt.Println("len:", len(ch), "cap:", cap(ch))
}
