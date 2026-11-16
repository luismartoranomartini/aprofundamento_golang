package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	a := s[0:3]
	b := s[2:5]
	a[2] = 99
	fmt.Println(s)
	fmt.Println(a, b)
	b[0] = 7
	fmt.Println(s, a)
}
