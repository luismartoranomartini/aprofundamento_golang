package main

import "fmt"

func main() {
	s := make([]int, 1000)
	pequeno := s[0:2]
	fmt.Println(len(pequeno), cap(pequeno))
}
