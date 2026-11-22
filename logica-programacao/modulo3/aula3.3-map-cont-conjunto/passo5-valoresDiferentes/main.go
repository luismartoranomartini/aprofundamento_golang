package main

import "fmt"

func contaDiferentes(s []int) int {
	conjunto := make(map[int]bool)

	for _, v := range s {
		conjunto[v] = true
		fmt.Println("li", v, "-> conjunto:", conjunto, "tamanho:", len(conjunto))
	}
	return len(conjunto)
}

func main() {
	fmt.Println("resposta:", contaDiferentes([]int{4, 4, 7, 4, 7, 9}))
}
