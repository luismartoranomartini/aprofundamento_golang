package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	Pessoa struct {
		Nome  string
		Idade int
	}
	Buscavel interface {
		constraints.Integer | ~string | Pessoa
	}
)

func Busca[T Buscavel](lista []T, f func(T) bool) int {
	for i, valor := range lista {
		if f(valor) {
			return i
		}
	}

	return -1
}

func main() {
	var (
		numeros  = []int{1, 2, 3, 4, 5}
		palavras = []string{"Ola", "Mundo"}
		pessoa   = []Pessoa{
			{
				Nome:  "Davi",
				Idade: 30,
			},
			{
				Nome:  "Luís",
				Idade: 49,
			},
			{
				Nome:  "Jonice",
				Idade: 59,
			},
		}
	)
	resultado := Busca(numeros, func(n int) bool {
		return n == 2
	})
	fmt.Println("Resultado dos números:", resultado)

	resultado2 := Busca(palavras, func(s string) bool {
		return len(s) == 4
	})
	fmt.Println("Posição da string:", resultado2)

	resultado3 := Busca(pessoa, func(p Pessoa) bool {
		return p.Idade == 59
		return p.Nome == "Luís"
	})
	fmt.Println("Posição da Pessoa:", resultado3)

}
