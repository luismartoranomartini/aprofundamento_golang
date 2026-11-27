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

func buscar(produtos []Produto, nome string) (Produto, bool) {
	for _, p := range produtos {
		if p.Nome == nome {
			return p, true
		}

	}
	return Produto{}, false
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

	produto = Produto{
		Nome:  "Caneta",
		Preco: 9.9,
		Qtd:   3,
	}
	produtos = acrescentar(produtos, produto)
	fmt.Println(produtos, len(produtos))

	busca, ok := buscar(produtos, "Lápis")
	fmt.Println(busca, ok)

	busca, ok = buscar(produtos, "Régua")
	fmt.Println(busca, ok)

	busca, ok = buscar(produtos, "Caneta")
	fmt.Println(busca, ok)

	busca, ok = buscar(nil, "Caneta")
	fmt.Println(busca, ok)
}
