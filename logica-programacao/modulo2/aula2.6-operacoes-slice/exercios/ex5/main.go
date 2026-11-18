package main

import "fmt"

// unicos devolve um slice NOVO com os elementos de s sem repetição,
// mantendo a ordem da primeira aparição de cada um. s não pode ser alterado.
// Regras: força bruta, sem map. Função auxiliar é permitida.

func existeRepetidos(s []int, v int) bool {
	for _, r := range s {
		if r == v {
			return true
		}
	}
	return false
}

func unicos(s []int) []int {
	var res []int
	for _, u := range s {
		if !existeRepetidos(res, u) {
			res = append(res, u)
		}
	}

	return res // TODO: escreva seu código
}

func main() {
	fmt.Println(unicos([]int{3, 1, 3, 2, 1, 2})) // esperado: [3 1 2]
	fmt.Println(unicos([]int{4, 4, 4}))          // esperado: [4]
	fmt.Println(unicos(nil))                     // esperado: []
}
