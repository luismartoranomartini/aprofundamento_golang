package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	t := s[1:3]
	fmt.Println(t, len(t), cap(t))
	t = t[:cap(t)]
	fmt.Println(t, len(t), cap(t))
}
