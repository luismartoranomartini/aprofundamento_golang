package main

import "fmt"

type Produto struct {
	Nome    string
	Preco   float64
	Estoque int
	Ativo   bool
}

func main() {
	p := Produto{Nome: "Caneta", Preco: 2.5}
	fmt.Printf("%+v\n", p)

	fmt.Println("preco", p.Preco)

	p.Estoque = 40
	p.Ativo = true
	fmt.Printf("%+v", p)
}
