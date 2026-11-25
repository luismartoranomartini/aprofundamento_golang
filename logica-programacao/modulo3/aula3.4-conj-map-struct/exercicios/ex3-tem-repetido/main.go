package main

import "fmt"

func temRepetido(num []int) bool {
	conjuntos := make(map[int]struct{})

	for _, n := range num {
		fmt.Println(conjuntos)
		conjuntos[n] = struct{}{}

	}
	if len(num) > len(conjuntos) {
		return true
	}

	return false
}

func main() {
	entrada := []int{3, 1, 4, 1, 5}
	fmt.Println(temRepetido(entrada))

	entrada = []int{3, 1, 4}
	fmt.Println(temRepetido(entrada))
}
