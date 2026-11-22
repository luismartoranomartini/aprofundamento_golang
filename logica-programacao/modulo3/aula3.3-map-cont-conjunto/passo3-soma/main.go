package main

import "fmt"

func main() {
	produtos := []string{"pão", "leite", "pão", "ovo", "leite", "pão"}
	valores := []int{8, 6, 8, 12, 6, 8}

	total := make(map[string]int)

	for i := range produtos {
		total[produtos[i]] += valores[i]

		fmt.Println("venda", i, ":", produtos[i], valores[i], "->", total)
	}
	fmt.Println("total: ", total)
}
