package main

import "fmt"

func meio(s []int) []int {
	if len(s) < 3 {
		return []int{}
	}
	return s[1 : len(s)-1]
}

func main() {
	casosTeste := [][]int{
		{},
		{10},
		{10, 20},
		{10, 20, 30},
		{10, 20, 30, 40, 50},
	}

	for _, caso := range casosTeste {
		resultado := meio(caso)
		fmt.Printf("Entrada (len=%d): %-18v -> Saída: %v\n", len(caso), caso, resultado)
	}
}

/*
A função meio precisa validar se o slice tem pelo menos 3 elementos antes de fatiar (s[1 : len(s)-1]), evitando assim um panic por índice fora dos limites (slice bounds out of range) quando o slice for muito pequeno.
*/
