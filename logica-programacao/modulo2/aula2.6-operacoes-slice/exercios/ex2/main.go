package main

import "fmt"

// semValor devolve um slice NOVO com os elementos de s diferentes de v,
// na mesma ordem. s não pode ser alterado.
func semValor(s []int, v int) []int {
	var res []int
	for _, x := range s {
		if x != v {
			res = append(res, x)
		}

	}

	return res
}

func main() {
	x := []int{1, 2, 1, 3, 1}
	r := semValor(x, 1)
	fmt.Println(r, x) // esperado: [2 3] [1 2 1 3 1]

	fmt.Println(semValor(nil, 1), semValor([]int{5, 5}, 5))
	// esperado: [] []
}
