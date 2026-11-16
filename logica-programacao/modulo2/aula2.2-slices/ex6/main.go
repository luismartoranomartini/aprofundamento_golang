package main

import "fmt"

func dobrar(s []int) {
	for i := range s {
		s[i] *= 2
	}
	fmt.Println("dentro:", s)
}

func main() {
	s := []int{10, 20, 30}
	t := s
	t[0] = 0
	fmt.Println(s, t)
	dobrar(s)
	fmt.Println("fora: ", s)
}
