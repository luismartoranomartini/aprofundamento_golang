package main

import "fmt"

// clonaMatriz devolve uma cópia de m em que nenhuma linha compartilha array com m.
// Alterar a cópia não pode alterar m.
func clonaMatriz(m [][]int) [][]int {

	profunda := make([][]int, len(m))

	for i := range m {
		profunda[i] = make([]int, len(m[i]))
		copy(profunda[i], m[i])
	}

	return profunda // TODO: escreva seu código
}

func main() {
	m := [][]int{{1, 2}, {3, 4, 5}}
	c := clonaMatriz(m)
	c[0][0] = 99
	c[1] = append(c[1], 6)
	fmt.Println(m) // esperado: [[1 2] [3 4 5]]
	fmt.Println(c) // esperado: [[99 2] [3 4 5 6]]
}
