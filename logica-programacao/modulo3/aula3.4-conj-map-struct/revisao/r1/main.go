package main

import "fmt"

func contarMaiores(num map[string]int, limite int) int {
	var guarda []int
	valores := 0

	for _, n := range num {
		if n > limite {
			guarda = append(guarda, n)
			valores++

		}
	}
	return valores
}

func main() {
	entrada := map[string]int{"a": 3, "b": 7, "c": 9}
	limite := 5

	fmt.Println(contarMaiores(entrada, limite))
}
