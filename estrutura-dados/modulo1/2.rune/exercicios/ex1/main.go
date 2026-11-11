package main

import "fmt"

func main() {
	s := "coração"
	fmt.Println("Byte: ", len(s))
	fmt.Println("Rune: ", len([]rune(s)))
	fmt.Printf("Comparação entre Byte e Rune é %v\n", len(s) == len([]rune(s)))
}
