package main

import "fmt"

func numeroFaltante(numero []int, n int) []int {
	conjunto := make(map[int]struct{})
	var guarda []int

	for _, x := range numero {
		conjunto[x] = struct{}{}
	}

	for i := 1; i <= n; i++ {
		if _, ok := conjunto[i]; !ok {
			guarda = append(guarda, i)
		}
	}
	return guarda
}

func main() {
	entrada := []int{2, 4, 4, 6, 1}
	n := 6
	fmt.Println(numeroFaltante(entrada, n))
}
