package main

import "fmt"

func uniao(a, b []int) []int {
	entra := make(map[int]bool)
	var guarda []int

	for _, x := range a {
		if !entra[x] {
			entra[x] = true
			guarda = append(guarda, x)

		}
	}

	for _, x := range b {
		if !entra[x] {
			entra[x] = true
			guarda = append(guarda, x)
		}
	}
	return guarda
}

func main() {
	a := []int{1, 2, 2, 3}
	b := []int{3, 4, 4, 5}
	fmt.Println(uniao(a, b))
}
