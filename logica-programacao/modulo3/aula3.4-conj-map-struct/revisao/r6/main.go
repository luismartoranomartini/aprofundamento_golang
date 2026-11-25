package main

import "fmt"

func emUmDosDois(a, b []int) []int {
	emA := make(map[int]struct{})
	emB := make(map[int]struct{})
	var guarda []int

	for _, n := range b {
		emB[n] = struct{}{}
	}
	for _, n := range a {
		emA[n] = struct{}{}
	}
	for _, n := range a {
		if _, ok := emB[n]; !ok {
			emB[n] = struct{}{}
			guarda = append(guarda, n)
		}
	}
	for _, n := range b {
		if _, ok := emA[n]; !ok {
			emA[n] = struct{}{}
			guarda = append(guarda, n)
		}
	}
	return guarda
}

func main() {
	a := []int{1, 2, 3, 3}
	b := []int{3, 4, 4, 5}

	fmt.Println(emUmDosDois(a, b))
}
