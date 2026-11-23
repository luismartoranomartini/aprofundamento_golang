package main

import "fmt"

func somaMaiorNotaAluno(nomes []string, notas []int) (string, int) {
	soma := make(map[string]int)

	for i := range nomes {
		soma[nomes[i]] += notas[i]
	}

	var maiorNome string
	var maiorSoma int

	for i, n := range soma {
		if soma[i] > maiorSoma {
			maiorSoma = n
			maiorNome = i
		}
	}
	return maiorNome, maiorSoma

}

func main() {
	nomes := []string{"Leo", "Mia", "Leo", "Mia", "Zeca"}
	notas := []int{5, 8, 9, 2, 7}

	fmt.Println(somaMaiorNotaAluno(nomes, notas))

}
