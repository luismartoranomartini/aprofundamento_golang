package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func Busca[T comparable](lista []T, f func(T) bool) int {
	for i, valor := range lista {
		if f(valor) {
			return i
		}
	}
	return -1
}

func main() {
	var (
		numeros  = []int{1, 2, 3, 4}
		palavras = []string{"Olá", "mundo"}
	)
	resultado := Busca(palavras, func(s string) bool {
		return len(s) > 3
	})

	resultado1 := Busca(numeros, func(n int) bool {
		return n == 3
	})

	fmt.Println("String:", resultado)
	fmt.Println("Número:", resultado1)
}
