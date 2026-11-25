package main

import "fmt"

func contarDistintos(str []string) int {
	conjunto := make(map[string]struct{})

	for _, n := range str {
		conjunto[n] = struct{}{}
	}
	return len(conjunto)
}

func main() {
	entrada := []string{"go", "map", "go", "rune", "map"}

	fmt.Println(contarDistintos(entrada))
}
