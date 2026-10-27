package main

import "fmt"

func SomaInteiros(numeros []int) int {
	var soma = 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {

	var (
		numerosInteiros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		numerosFloat    = []float64{1.0, 2.0, 3.0}
	)
	fmt.Println(SomaInteiros(numerosInteiros))
	fmt.Println(SomaInteiros(numerosFloat))
}
