package main

import "fmt"

func contagemLetras(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}

func main() {
	palavra := "bbcca"
	frequencia := contagemLetras(palavra)

	fmt.Println("Contagem de todas as letras:")
	for r, qtd := range frequencia {
		fmt.Printf("Letras '%c': %d vez(es)\n", r, qtd)
	}
}
