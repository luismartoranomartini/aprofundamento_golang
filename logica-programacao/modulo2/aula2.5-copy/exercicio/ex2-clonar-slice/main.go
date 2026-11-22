package main

import "fmt"

func clona(s []int) []int {
	dst := make([]int, len(s))
	copy(dst, s)
	return dst
}

func main() {
	s := []int{1, 2, 3}
	c := clona(s)
	c[0] = 100
	fmt.Println(s, c)
	fmt.Println(clona(nil), len(clona(nil)))
}
