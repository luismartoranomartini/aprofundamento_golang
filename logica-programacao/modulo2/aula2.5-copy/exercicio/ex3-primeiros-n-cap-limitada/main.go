package main

import "fmt"

func primeiros(s []int, n int) []int {
	if n < 0 || n > len(s) {
		return nil
	}

	return s[0:n:n]
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	r := primeiros(x, 2)
	r = append(r, 99)
	fmt.Println(x, r) // esperado: [1 2 3 4 5] [1 2 99]

	fmt.Println(primeiros(x, 0), primeiros(x, 6), primeiros(x, -1))
	// esperado: [] [] []
}
