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
		palavras = []string{"Olá", "Mundo"}
		pessoas  = []Pessoa{
			{
				Nome:  "Davi",
				Idade: 30,
			},
			{
				Nome:  "Lucas",
				Idade: 25,
			},
			{
				Nome:  "Pedro",
				Idade: 28,
			},
		}
	)
	resultado := Busca(numeros, func(n int) bool {
		return n == 3
	})
	fmt.Println(resultado)

	resultado2 := Busca(palavras, func(p string) bool {
		return p == "Olá"
	})

	fmt.Println(resultado2)

	resultado3 := Busca(pessoas, func(p Pessoa) bool {
		return p.Idade > 28 || p.Nome == "Davi"
	})
	fmt.Println(resultado3)

}
