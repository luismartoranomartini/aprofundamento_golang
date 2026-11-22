package main

import "fmt"

func soma(janela []int) int {
	total := 0
	for _, v := range janela {
		total += v
	}
	return total
}

func somasPrealocado(s []int, k int) []int {
	if k <= 0 || k > len(s) {
		return nil
	}

	res := make([]int, 0, len(s)-k+1)

	for i := 0; i < len(s)-k+1; i++ {
		res = append(res, soma(s[i:i+k]))
	}
	return res
}

func main() {
	seq := []int{2, 1, 5, 1, 3, 2}
	r := somasPrealocado(seq, 3)
	fmt.Println(r, len(r), cap(r)) // esperado: [8 7 9 6] 4 4

	fmt.Println(somasPrealocado(seq, 6), somasPrealocado(seq, 0), somasPrealocado(seq, 7))
	// esperado: [14] [] []
}
