package main

import "fmt"

func maiorTamanho(s []string) (int, []string) {
	grupos := make(map[int][]string)
	var quantidade int
	var palavras []string

	for _, n := range s {
		qtd := len(n)
		grupos[qtd] = append(grupos[qtd], n)
	}

	for i, n := range grupos {
		if i > quantidade {
			quantidade = i
			palavras = n

		}
	}

	return quantidade, palavras

}

func main() {
	entrada := []string{"a", "bola", "ar", "casa", "sol", "pe"}

	fmt.Println(maiorTamanho(entrada))
}
