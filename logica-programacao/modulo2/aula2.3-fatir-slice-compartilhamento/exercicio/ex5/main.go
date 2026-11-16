package main

import "fmt"

func soma(janela []int) int {
	somatorio := 0
	for _, j := range janela {
		somatorio += j
	}
	return somatorio
}

func maiorJanelas(s []int, k int) (maior int) {
	if k <= 0 || k > len(s) {
		return -1
	}
	var somarJanela int

	for i := 0; i <= len(s)-k; i++ {
		guarda := s[i : i+k]
		somarJanela = soma(guarda)
		if i == 0 || somarJanela > maior {
			maior = somarJanela
		}
	}
	return maior
}

func main() {
	sequencia := []int{2, 1, 5, 1, 3, 2}
	k := 3
	sequencia1 := []int{-3, -1, -2}
	c := 2

	r1 := maiorJanelas(sequencia, k)
	r2 := maiorJanelas(sequencia1, c)
	fmt.Println(r1)
	fmt.Println(r2)
}
