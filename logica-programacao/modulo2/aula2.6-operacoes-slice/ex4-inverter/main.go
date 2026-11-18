package main

import "fmt"

func inveter(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	par := []int{1, 2, 3, 4}
	inveter(par)
	fmt.Println("4 elementos:", par)

	impar := []int{1, 2, 3, 4, 5}
	inveter(impar)
	fmt.Println("5 elementos:", impar)

	var vazio []int
	inveter(vazio)
	fmt.Println("vazio:", vazio)

	um := []int{7}
	inveter(um)
	fmt.Println("1 elemento:", um)

	orig := []int{1, 2, 3}
	c := make([]int, len(orig))
	copy(c, orig)
	inveter(c)
	fmt.Println("original:", orig, "| invertido", c, "mesmo array?", orig[0] == c[0])
}
