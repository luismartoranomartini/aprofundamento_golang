package main

import "fmt"

func insereOrdenado(s []int, v int) []int {

	i := 0
	for i < len(s) && v >= s[i] {
		i++
	}
	s = append(s, v)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}

func main() {
	fmt.Println(insereOrdenado([]int{1, 3, 5, 7}, 4)) // esperado: [1 3 4 5 7]
	fmt.Println(insereOrdenado([]int{1, 3, 5, 7}, 0)) // esperado: [0 1 3 5 7]
	fmt.Println(insereOrdenado([]int{1, 3, 5, 7}, 9)) // esperado: [1 3 5 7 9]
	fmt.Println(insereOrdenado(nil, 2))               // esperado: [2]
}
