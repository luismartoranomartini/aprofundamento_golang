package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
	Qtd   int
}

func acrescentar(produtos []Produto, produto Produto) []Produto {
	return append(produtos, produto)
}

func main() {
	var produtos []Produto

	produto := Produto{
		Nome:  "Caneta",
		Preco: 2.5,
		Qtd:   10,
	}
	produtos = acrescentar(produtos, produto)
	fmt.Println(produtos, len(produtos))

	produto = Produto{
		Nome:  "Lápis",
		Preco: 1.2,
		Qtd:   50,
	}
	produtos = acrescentar(produtos, produto)
	fmt.Println(produtos, len(produtos))
}
