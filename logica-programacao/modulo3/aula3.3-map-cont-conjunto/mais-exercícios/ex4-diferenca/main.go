package main

import "fmt"

func diferenca(a, b []int) []int {
	jaEntrou := make(map[int]bool)
	unicos := make(map[int]bool)
	var guard []int

	for _, x := range b {
		unicos[x] = true
	}

	for _, x := range a {
		if !unicos[x] && !jaEntrou[x] {
			guard = append(guard, x)
			jaEntrou[x] = true
		}
	}
	return guard
}

func main() {
	// a := []int{1, 2, 2, 3, 4}
	a := []int{1, 1, 3}
	b := []int{2}
	// b := []int{2, 4, 4, 5}
	fmt.Println(diferenca(a, b))
}
