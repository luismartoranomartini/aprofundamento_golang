package main

import "fmt"

func notaAluno(nomes []string, notas []int) map[string][]int {
	resultado := make(map[string][]int)

	for i, n := range nomes {
		resultado[n] = append(resultado[n], notas[i])
	}
	return resultado
}

func main() {
	nomes := []string{"Leo", "Mia", "Leo", "Mia", "Zeca"}
	notas := []int{5, 8, 9, 2, 7}

	fmt.Println(notaAluno(nomes, notas))
}
