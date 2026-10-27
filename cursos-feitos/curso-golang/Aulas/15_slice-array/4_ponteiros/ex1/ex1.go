package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	var c *int
	fmt.Println("Variavel a:", a)
	fmt.Println("Variavel a (endereço):", &a)
	fmt.Println("Variavel b (endereço):", b)
	fmt.Println("Variável b:", *b)
	fmt.Println("Variável c:", c) // nil
}
