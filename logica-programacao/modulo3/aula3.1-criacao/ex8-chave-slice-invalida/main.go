package main

import "fmt"

func main() {
	// A chave precisa ser de um tipo que aceite ==.
	// Slice não aceita ==, então não pode ser chave.
	m := map[[]int]string{}
	fmt.Println(m)

	/*
		# command-line-arguments
		./main.go:8:11: invalid map key type []int
	*/
}
