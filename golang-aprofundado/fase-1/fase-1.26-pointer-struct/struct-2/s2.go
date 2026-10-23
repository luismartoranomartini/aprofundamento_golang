package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	a := Pessoa{"Ana", 30}
	p := &a

	fmt.Printf("tipo de p: %T\n", p)
	fmt.Println(p.Idade, (*p).Idade)
	p.Idade = 31
	fmt.Println("a:", a)

	q := &Pessoa{"Beto", 20}
	q.Idade++
	fmt.Println(*q)

	n := new(Pessoa)
	fmt.Printf("%+v\n", *n)

	pi := &a.Idade
	*pi = 50
	fmt.Println("a:", a)
}
