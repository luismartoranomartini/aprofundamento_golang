package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	t := s[1:3]
	fmt.Println(t[:4])
	fmt.Println(t[:5])
}
