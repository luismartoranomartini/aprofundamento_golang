package main

import "fmt"

func pertence(s []int, n int) bool {
	conjunto := make(map[int]struct{})

	for _, x := range s {
		conjunto[x] = struct{}{}
	}

	_, ok := conjunto[n]

	return ok
}

func main() {
	num := []int{4, 8, 15, 16}
	n := 15

	fmt.Println(pertence(num, n))

	n1 := 3

	fmt.Println(pertence(num, n1))
}
