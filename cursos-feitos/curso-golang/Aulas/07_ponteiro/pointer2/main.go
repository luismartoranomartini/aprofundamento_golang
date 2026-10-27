package main

import "fmt"

func incrementaErrado(n int) {
	n++
}

func incrementaCorreto(n *int) {
	*n++
}

func main() {
	x := 5
	incrementaErrado(x)
	fmt.Println(x)

	incrementaCorreto(&x)
	fmt.Println(x)
}
