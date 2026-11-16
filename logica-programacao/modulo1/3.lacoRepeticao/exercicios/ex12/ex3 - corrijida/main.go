package main

import "fmt"

func main() {
	i := 0
	x := i // x recebe 0
	i++    // i passa a valer 1

	fmt.Println("x:", x, "| i:", i) // Saída: x: 0 | i: 1
}
