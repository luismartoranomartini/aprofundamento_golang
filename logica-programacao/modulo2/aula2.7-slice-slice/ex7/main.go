package main

import "fmt"

func main() {
	var m [][]int
	fmt.Println(m == nil, len(m))

	for _, linha := range m {
		fmt.Println(linha)
	}

	m = append(m, []int{})
	fmt.Println(m, len(m), len(m[0]))
}
