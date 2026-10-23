package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	var z Pessoa
	fmt.Printf("%q %d\n", z.Nome, z.Idade)

	a := Pessoa{
		Nome:  "Ana",
		Idade: 30,
	}
	fmt.Println(a)

	b := a
	b.Idade = 99
	fmt.Println("a:", a, "| b:", b)

	fmt.Println("a == b:", a == b)
	b.Idade = 30
	fmt.Println("a == b:", a == b)
}
