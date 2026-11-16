package main

import "fmt"

func realocacoes(s []int, n int) int {
	mudancas := 0
	for i := 0; i < n; i++ {
		antes := cap(s)
		s = append(s, i)
		if cap(s) != antes {
			mudancas++
		}
	}
	return mudancas
}

func main() {
	n := 1000

	var a []int
	fmt.Println("sem reservar:     ", realocacoes(a, n), "realocaçoes")

	b := make([]int, 0, n)
	fmt.Println("com make([]int, 0, n)", realocacoes(b, n), "realocaçoes")
}
