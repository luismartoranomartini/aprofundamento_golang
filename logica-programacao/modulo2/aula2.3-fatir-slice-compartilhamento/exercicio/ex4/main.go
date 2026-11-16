package main

import "fmt"

func maiorJanela(s []int, k int) int {
	if k <= 0 || k > len(s) {
		return -1
	}
	var maxSoma int

	for i := 0; i <= len(s)-k; i++ {
		janela := s[i : i+k]
		somaAtual := 0

		for _, valor := range janela {
			somaAtual += valor
		}

		if i == 0 || somaAtual > maxSoma {
			maxSoma = somaAtual
		}
	}
	return maxSoma
}

func main() {
	numeros := []int{2, 1, 5, 1, 3, 2}
	fmt.Println(maiorJanela(numeros, 3))
	fmt.Println(maiorJanela(numeros, 0))
	fmt.Println(maiorJanela(numeros, 7))
}

/*
CONFESSO: pedi para o Gemini fazer. Achei muito difícil
*/
