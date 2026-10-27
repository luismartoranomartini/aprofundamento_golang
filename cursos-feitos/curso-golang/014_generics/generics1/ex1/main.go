package main

import "fmt"

func SomaInteiros(numeros []int) int {
	var soma = 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaFloat(numeros []float64) float64 {
	var soma float64 = 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaAny(numeros []any) any {
	var soma float64 = 0
	for _, numero := range numeros {
		switch v := numero.(type) {
		case float64:
			soma += v
		case int:
			soma += float64(v)
		}
	}
	return soma
}

func SomaGenerics[T int | float64 | uint](numeros []T) T {
	var soma T = 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {
	var (
		numerosInteiros = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		numeroFloat     = []float64{1.0, 2.0, 3.0}
		numeroUint      = []uint{2, 4, 6, 8, 12}
	)
	fmt.Println("-->Sem Generics<--")
	fmt.Println(SomaInteiros(numerosInteiros))
	fmt.Println(SomaFloat(numeroFloat))

	fmt.Println("-->Com Generics<--")
	fmt.Println(SomaGenerics(numerosInteiros))
	fmt.Println(SomaGenerics(numeroFloat))
	fmt.Println(SomaGenerics(numeroUint))
}
