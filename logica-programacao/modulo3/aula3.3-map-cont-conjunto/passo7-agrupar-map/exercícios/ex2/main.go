package main

import "fmt"

func agruparPalavras(s []string) map[int][]string {
	grupos := make(map[int][]string)

	for _, n := range s {
		qtd := len(n)
		grupos[qtd] = append(grupos[qtd], n)
	}

	return grupos
}

func main() {
	entrada := []string{"a", "bola", "ar", "casa", "sol", "pe"}

	fmt.Println(agruparPalavras(entrada))
}
