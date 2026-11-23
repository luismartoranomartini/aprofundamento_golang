package main

import "fmt"

func interseccao(a, b []int) []int {
	unicos := make(map[int]bool)
	var guarda []int

	for _, x := range b {
		unicos[x] = true
	}

	for _, x := range a {
		if unicos[x] {
			guarda = append(guarda, x)
			unicos[x] = false
		}
	}

	return guarda
}

func main() {
	a := []int{1, 2, 2, 3, 4}
	b := []int{2, 4, 4, 5}
	fmt.Println(interseccao(a, b)) // [2 4]
}
