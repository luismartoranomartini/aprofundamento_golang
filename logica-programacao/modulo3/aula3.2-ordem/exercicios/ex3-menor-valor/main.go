package main

import "fmt"

func menorValor(m map[string]int) (string, bool) {
	if len(m) == 0 {
		return "", false
	}
	var menorChave string
	var menorValor int
	primeiro := true

	for k, v := range m {
		if primeiro || v < menorValor {
			menorChave = k
			menorValor = v
			primeiro = false
		}
	}
	return menorChave, true
}

func main() {
	mapa := map[string]int{"a": 5, "b": 2, "c": 9}

	chave, ok := menorValor(mapa)
	fmt.Println(chave, ok)

	mapa1 := map[string]int{"a": -5, "b": -2, "c": -9}

	chave1, ok := menorValor(mapa1)
	fmt.Println(chave1, ok)

	mapa2 := map[string]int{}

	chave2, ok := menorValor(mapa2)
	fmt.Println(chave2, ok)

}
