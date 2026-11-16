package main

import "fmt"

func main() {
	s := make([]int, 5)
	for i := range s {
		n := i + 1
		s[i] = n * n
	}
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))
}
