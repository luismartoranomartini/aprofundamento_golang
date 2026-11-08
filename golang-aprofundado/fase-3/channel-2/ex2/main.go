package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case v := <-ch:
		fmt.Println("recebi", v)
	default:
		fmt.Println("nada pronto, sigo em frente")
	}

	select {
	case ch <- 5:
		fmt.Println("enviei 5")
	default:
		fmt.Println("canal cheio")
	}
	select {
	case ch <- 6:
		fmt.Println("enviei 6")
	default:
		fmt.Println("canal cheio, descartei o 6")
	}
	fmt.Println("no canal", <-ch)
}
