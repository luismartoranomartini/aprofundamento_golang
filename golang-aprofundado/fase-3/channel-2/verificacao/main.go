package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case ch <- 1:
		fmt.Println("A")
	default:
		fmt.Println("B")
	}

	select {
	case ch <- 2:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	fmt.Println(<-ch)
}
