package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func aniversarioV(p Pessoa) {
	p.Idade++
}

func aniversarioP(p *Pessoa) {
	p.Idade++
}

func main() {
	a := Pessoa{"Ana", 30}
	aniversarioV(a)
	fmt.Println("depois de V:", a.Idade)
	aniversarioP(&a)
	fmt.Println("depois de P:", a.Idade)

	var np *Pessoa
	defer func() {
		fmt.Println("recuperado:", recover())
	}()
	fmt.Println(np.Idade)
}
