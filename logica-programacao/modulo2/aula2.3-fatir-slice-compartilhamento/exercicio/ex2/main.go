package main

import "fmt"

func main() {
	s := make([]int, 3, 6)
	t := s[:5]
	fmt.Println(t, len(t), cap(t))
	fmt.Println(s[:7])
}
