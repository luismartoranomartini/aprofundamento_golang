package main

import "fmt"

func main() {
	var estoque int

	estoque = 20

	if estoque > 0 {
		fmt.Println("Produto disponível")
	} else {
		fmt.Println("produto esgotado")
	}
}
