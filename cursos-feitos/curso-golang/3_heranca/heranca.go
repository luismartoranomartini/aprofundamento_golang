package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

type estudante struct {
	pessoa    pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{
		nome:      "João",
		sobrenome: "Silva",
		idade:     19,
		altura:    178,
	}
	fmt.Println(p1)

	// e1 := estudante{p1, "USP", "Engenharia"}
	// fmt.Println(e1.nome)

	e2 := estudante{
		pessoa:    p1,
		faculdade: "USP",
		curso:     "Engenharia",
	}
	fmt.Println(e2)
}
