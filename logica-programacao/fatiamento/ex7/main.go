package main

import (
	"fmt"
	"sort"
)

type Par struct {
	Letra rune
	Qtd   int
}

func main() {
	var palavra string
	fmt.Println("Digite uma palavra: ")
	fmt.Scan(&palavra)

	m := make(map[rune]int)
	for _, r := range palavra {
		m[r]++
	}
	var lista []Par
	for r, qtd := range m {
		lista = append(lista, Par{Letra: r, Qtd: qtd})
	}

	sort.Slice(lista, func(i, j int) bool {
		return lista[i].Letra < lista[j].Letra
	})

	fmt.Println("Ordem alfabética")
	for _, p := range lista {
		fmt.Printf("Letra '%c': %d vez(es)\n", p.Letra, p.Qtd)
	}
}
