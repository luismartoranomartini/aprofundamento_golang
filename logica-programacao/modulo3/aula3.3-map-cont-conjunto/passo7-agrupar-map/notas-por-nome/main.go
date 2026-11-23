package main

import "fmt"

func main() {
	nomes := []string{"Ana", "Bia", "Ana", "Caio", "Bia", "Ana"}
	notas := []int{10, 7, 4, 8, 9, 6}

	porNome := make(map[string][]int)

	for i := 0; i < len(nomes); i++ {
		porNome[nomes[i]] = append(porNome[nomes[i]], notas[i])
	}

	ordem := []string{"Ana", "Bia", "Caio"}
	for _, n := range ordem {
		fmt.Println(n, porNome[n])
	}

	for _, n := range ordem {
		soma := 0
		for _, v := range porNome[n] {
			soma += v
		}
		fmt.Println(n, "soma:", soma)
	}
}
