package main

import "fmt"

func uniao(a, b []int) []int {
	mapa := make(map[int]bool)
	var guarda []int
	var resultado []int

	guarda = append(guarda, a...)
	guarda = append(guarda, b...)
	for _, x := range guarda {
		if !mapa[x] {
			mapa[x] = true
			resultado = append(resultado, x)
		}
	}
	return resultado
}

func main() {
	fmt.Println(uniao([]int{1, 2, 2, 3}, []int{3, 4, 4, 5})) // [1 2 3 4 5]
	fmt.Println(uniao([]int{}, []int{2, 2, 1}))              // [2 1]
	fmt.Println(uniao([]int{7, 7}, nil))                     // [7]
	fmt.Println(uniao(nil, nil))
}
