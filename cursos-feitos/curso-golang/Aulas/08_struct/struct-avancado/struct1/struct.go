package main

import "fmt"

type Produto struct {
	Nome     string
	Preco    float64
	Desconto float64
}

// Método receiver
func (p Produto) PrecoComDesconto() float64 {
	return p.Preco * (1 - p.Desconto)
}

func main() {
	var produto1 Produto
	produto1 = Produto{
		Nome:     "Lápis",
		Preco:    1.79,
		Desconto: 0.05,
	}
	fmt.Println(produto1, produto1.PrecoComDesconto())

	produto2 := Produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2, produto2.PrecoComDesconto())
}
