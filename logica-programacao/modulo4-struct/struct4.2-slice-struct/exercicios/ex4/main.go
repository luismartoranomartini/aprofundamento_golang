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

func estoqueBaixo(produtos []Produto, limite int) []Produto {
	var resultado []Produto
	for _, p := range produtos {
		if p.Qtd < limite {
			resultado = append(resultado, p)
		}
	}
	return resultado
}

func reporEstoque(produtos []Produto, nome string, qtd int) bool {
	for _, p := range produtos {
		if p.Nome == nome {
			p.Qtd += qtd
			return true
		}
	}

	return false
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

	estoque := estoqueBaixo(produtos, 20)
	fmt.Println(estoque, len(estoque))

	estoque = estoqueBaixo(produtos, 1)
	fmt.Println(estoque, len(estoque))

	estoque = estoqueBaixo(produtos, 100)
	fmt.Println(estoque, len(estoque))

	repor := reporEstoque(produtos, "Lápis", 5)
	fmt.Println(repor, produtos)

	repor = reporEstoque(produtos, "Régua", 5)
	fmt.Println(repor, produtos)
}
