package main

import "fmt"

type Endereco struct {
	Rua    string
	Cidade string
}

type Pessoa struct {
	Nome string
	End  Endereco
}

func main() {
	p := Pessoa{
		Nome: "Ana",
		End:  Endereco{Rua: "Rua A", Cidade: "Recife"},
	}

	fmt.Println(p.End.Cidade)
	p.End.Cidade = "Olinda"
	fmt.Printf("%+v", p)
}
