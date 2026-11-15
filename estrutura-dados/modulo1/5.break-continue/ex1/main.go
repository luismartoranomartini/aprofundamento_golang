package main

import "fmt"

func main() {
	s := "abc7d94dc81"
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			fmt.Printf("Posição dos caracteres: %d\n", i)
			fmt.Printf("O digito que está na posição %d é %c\n", i, s[i])
			count++
			break

		}
	}
	fmt.Println("total de digitos: ", count)
	fmt.Println("Depois do laço")
}
