package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	s := []string{"x", "y", "z"}

	for i, v := range slices.All(s) {
		fmt.Print(i, v, " ")
	}
	fmt.Println()
	for v := range slices.Values(s) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	for i, v := range slices.Backward(s) {
		fmt.Print(i, v, " ")
	}
	fmt.Println()

	m := map[string]int{"b": 2, "a": 1, "c": 3}
	chaves := slices.Sorted(maps.Keys(m))
	fmt.Println(chaves)
	fmt.Println(slices.Collect(maps.Values(map[string]int{"k": 7})))
}
