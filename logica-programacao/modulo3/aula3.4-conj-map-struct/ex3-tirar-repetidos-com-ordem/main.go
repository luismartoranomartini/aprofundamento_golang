package main

import "fmt"

func main() {
	entrada := []int{3, 1, 3, 2, 1, 4}

	visto := make(map[int]struct{})
	var unicos []int

	for _, n := range entrada {
		if _, ja := visto[n]; !ja {
			visto[n] = struct{}{}
			unicos = append(unicos, n)
		}
	}
	fmt.Println(entrada, "->", unicos)
}
