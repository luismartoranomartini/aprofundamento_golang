package main

import "fmt"

type Caixa struct {
	Rotulo string
	Itens  []int
}

func main() {
	a := Caixa{"A", []int{1, 2, 3}}
	b := a
	b.Rotulo = "B"
	b.Itens[0] = 100
	b.Itens = append(b.Itens, 4)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
