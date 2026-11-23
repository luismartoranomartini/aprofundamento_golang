package main

import "fmt"

func diferenca(a, b []int) []int {
	ignora := make(map[int]bool)
	var guarda []int

	for _, x := range b { // entra com true
		ignora[x] = true
	}

	for _, x := range a {
		if !ignora[x] {
			guarda = append(guarda, x)
			ignora[x] = true
		}
	}
	return guarda
}

func main() {
	// a := []int{1, 2, 2, 3, 4}
	a := []int{1, 1, 3}
	b := []int{2}
	// b := []int{2, 4, 4, 5}
	fmt.Println(diferenca(a, b))
}
