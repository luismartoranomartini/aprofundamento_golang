package main

import (
	"fmt"
)

// triangulo devolve n linhas: a linha i tem i+1 elementos, de 1 até i+1.
func triangulo(n int) [][]int {

	tri := make([][]int, n)
	for i := range tri {
		tri[i] = make([]int, i+1)
		for j := range tri[i] {
			tri[i][j] = j + 1
		}
	}

	return tri // TODO: escreva seu código
}

func main() {
	fmt.Println(triangulo(3)) // esperado: [[1] [1 2] [1 2 3]]
	fmt.Println(triangulo(1)) // esperado: [[1]]
	fmt.Println(triangulo(0)) // esperado: []
}
