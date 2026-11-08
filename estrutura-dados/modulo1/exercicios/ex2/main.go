package main

import "fmt"

func main() {
	palavra := "Gópher"

	// entrega byte (número)
	for i := 0; i < len(palavra); i++ {
		saida := palavra[i]
		fmt.Printf("Índice: %d | Byte: %d | Caractere: %c\n", i, saida, saida)
	}

	// entrega rune
	for i, p := range palavra {
		fmt.Printf("Índice: %d | Byte: %d | Caractere: %c\n", i, p, p)
	}
}
