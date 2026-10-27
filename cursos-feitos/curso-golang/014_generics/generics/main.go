package main

import "fmt"

func SomaGenerics[T int | float64](numeros []T) T {
	var soma T = 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {
	numerosInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numerosFloat := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println(SomaGenerics(numerosInt))
	fmt.Println(SomaGenerics(numerosFloat))

}
