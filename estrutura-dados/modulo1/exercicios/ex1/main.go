package main

import "fmt"

func main() {
	s := "Linguagem Go"
	fmt.Println(len(s))
	ultimo := s[len(s)-1]
	fmt.Printf("Primeiro byte: %d/caractere %c, último byte: %d/caractere %c\n", s[0], s[0], ultimo, ultimo)
}
