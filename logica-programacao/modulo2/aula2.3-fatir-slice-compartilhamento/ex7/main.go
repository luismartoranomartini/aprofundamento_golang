package main

import "fmt"

func main() {
	var m []int
	fmt.Println(m == nil, len(m), cap(m))
	s := []int{1, 2, 3}
	v := s[1:1]
	fmt.Println(v == nil, len(v), cap(v))
}
