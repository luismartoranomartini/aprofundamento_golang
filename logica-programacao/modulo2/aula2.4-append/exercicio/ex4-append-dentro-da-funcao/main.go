package main

import "fmt"

func adiciona(s []int, v int) {
	s = append(s, v)
	s[0] = 100
}

func main() {
	a := make([]int, 1, 1)
	b := make([]int, 1, 4)
	adiciona(a, 7)
	adiciona(b, 7)
	fmt.Println(a, b, b[:2])
}
