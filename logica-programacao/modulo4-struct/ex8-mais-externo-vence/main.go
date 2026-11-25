package main

import "fmt"

type Casa struct {
	Nome   string
	Cidade string
}

type Dono struct {
	Nome string
	Casa
}

func main() {
	d := Dono{Nome: "Ana", Casa: Casa{Nome: "Sobrado", Cidade: "Recife"}}

	fmt.Println(d.Nome)
	fmt.Println(d.Casa.Nome)
	fmt.Println(d.Cidade)
}
