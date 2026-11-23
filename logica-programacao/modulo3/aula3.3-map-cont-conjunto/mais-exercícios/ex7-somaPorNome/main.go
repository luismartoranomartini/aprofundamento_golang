package main

import "fmt"

func somaPorNomes(nomes []string, valores []int) map[string]int {
	contador := make(map[string]int)

	for i := range nomes {
		contador[nomes[i]] += valores[i]
	}

	return contador
}

func main() {
	nomes := []string{"Ana", "Bia", "Ana"}
	valores := []int{10, 5, 7}

	fmt.Println(somaPorNomes(nomes, valores))
}
