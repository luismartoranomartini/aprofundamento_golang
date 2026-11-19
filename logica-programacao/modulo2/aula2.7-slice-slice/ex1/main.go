package main

import "fmt"

func main() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(len(m[0]))
	fmt.Println(m[1][2])

	for i, linha := range m {
		for j, v := range linha {
			fmt.Println(i, j, v)
		}
	}
}
