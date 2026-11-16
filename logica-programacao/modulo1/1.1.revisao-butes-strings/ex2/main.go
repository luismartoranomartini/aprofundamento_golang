package main

import "fmt"

func main() {
	s := "Go!"
	fmt.Printf("Caracter: %c | Número %d\n", s[len(s)-1], s[len(s)-1])
}
