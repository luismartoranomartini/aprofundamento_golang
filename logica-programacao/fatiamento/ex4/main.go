package main

import "fmt"

func primeiroNaoRepetido(s string) (rune, bool) {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	for _, r := range s {
		if m[r] == 1 {
			return r, true
		}
	}
	return 0, false
}

func main() {
	palavra := "bbcca"
	if r, ok := primeiroNaoRepetido(palavra); ok {
		fmt.Printf("A primeira letra que não se repete é: %c\n", r)
	} else {
		fmt.Println("Não há caracteres únicos na palavra.")
	}
}
