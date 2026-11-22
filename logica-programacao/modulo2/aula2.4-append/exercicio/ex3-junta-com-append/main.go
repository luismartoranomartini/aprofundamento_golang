package main

import "fmt"

func junta(a, b []int) []int {
	return append(a, b...)
}

func main() {
	// a, b := []int{1, 2}, []int{3, 4, 5}
	// r := junta(a, b)
	// r[0] = 99
	// fmt.Println(a, b, r)

	// Caso 1
	a, b := []int{1, 2}, []int{3, 4, 5}
	r := junta(a, b)
	r[0] = 99
	fmt.Println(a, b, r) // esperado: [1 2] [3 4 5] [99 2 3 4 5]

	// Caso 2: a2 tem len 2 e cap 10
	a2 := make([]int, 2, 10)
	a2[0], a2[1] = 1, 2
	r2 := junta(a2, b)
	r2[0] = 77
	fmt.Println(a2, r2) // esperado: [1 2] [77 2 3 4 5]
}
