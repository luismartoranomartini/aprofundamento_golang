package main

import "fmt"

func somaLinhas(m [][]int) []int {
	var acumula []int
	var soma int

	for i := 0; i < len(m); i++ {
		soma = 0
		for j := 0; j < len(m[i]); j++ {
			soma += m[i][j]
		}
		acumula = append(acumula, soma)
	}

	return acumula
}

func main() {

	a := [][]int{{1, 2, 3}, {4, 5}, {6}}
	fmt.Println(somaLinhas(a)) // esperado: [6 9 6]

	b := [][]int{{}, {7, 8}}
	fmt.Println(somaLinhas(b)) // esperado: [0 15]

	fmt.Println(somaLinhas(nil)) // esperado: []

}
