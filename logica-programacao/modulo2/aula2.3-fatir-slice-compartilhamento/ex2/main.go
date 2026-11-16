package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	t := s[1:4]
	u := t[1:3]
	fmt.Println(t, u)
	fmt.Println(t[0], u[0])
}
