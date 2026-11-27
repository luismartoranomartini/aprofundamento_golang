package main

import "fmt"

type Venda struct {
	Vendedor string
	Produto  string
	Valor    float64
}

func vendaMaior(vendas []Venda) (string, float64) {
	contagem := make(map[string]float64)
	var melhorVendedor string
	var maiorQtd float64

	for _, v := range vendas {
		contagem[v.Vendedor] += v.Valor
	}
	for vendedor, total := range contagem {
		if total > maiorQtd {
			maiorQtd = total
			melhorVendedor = vendedor
		}
	}
	return melhorVendedor, maiorQtd
}

func main() {
	vendedor := []Venda{
		{Vendedor: "Ana", Produto: "Teclado", Valor: 150},
		{Vendedor: "Leo", Produto: "Monitor", Valor: 800},
		{Vendedor: "Ana", Produto: "Mouse", Valor: 100},
		{Vendedor: "Leo", Produto: "Cabo", Valor: 50},
		{Vendedor: "Ana", Produto: "Cadeira", Valor: 700},
	}

	maiorVendedor, melhorVenda := vendaMaior(vendedor)
	fmt.Println(maiorVendedor, melhorVenda)
}
