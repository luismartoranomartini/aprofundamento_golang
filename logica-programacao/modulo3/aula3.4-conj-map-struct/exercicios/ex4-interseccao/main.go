package main

import "fmt"

func interseccao(a, b []int) []int {
	unicos := make(map[int]struct{})
	var guarda []int

	for _, n := range b {
		unicos[n] = struct{}{}
	}
	for _, n := range a {
		if _, ok := unicos[n]; ok {
			delete(unicos, n)
			guarda = append(guarda, n)
		}
	}
	return guarda
}

func main() {
	a := []int{1, 2, 2, 3, 4}
	b := []int{2, 3, 3, 5}
	fmt.Println(interseccao(a, b))
}
