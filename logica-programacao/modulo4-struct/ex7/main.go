package main

import "fmt"

type Endereco struct {
	Rua    string
	Cidade string
}

type PessoaA struct {
	Nome string
	End  Endereco
}

type PessoaB struct {
	Nome string
	Endereco
}

func main() {
	a := PessoaA{Nome: "Ana", End: Endereco{Rua: "Rua A", Cidade: "Recife"}}
	b := PessoaB{Nome: "Leo", Endereco: Endereco{Rua: "Rua B", Cidade: "Recife"}}

	fmt.Println(a.End.Cidade)
	fmt.Println(b.Endereco.Cidade)
	fmt.Println(b.Cidade)

	b.Cidade = "Olinda"
	fmt.Println(b.Endereco.Cidade)
	fmt.Println(b.Cidade)

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)

}
