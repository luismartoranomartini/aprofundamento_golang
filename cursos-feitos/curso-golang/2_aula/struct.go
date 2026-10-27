package main

import "fmt"

type Usuario struct {
	nome     string
	idade    uint
	Endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     uint
}

func main() {
	fmt.Println("Arquivo structs")
	endereco := Endereco{
		logradouro: "Aqui",
		numero:     0,
	}

	u := Usuario{
		nome:  "Luís",
		idade: 49,
	}

	fmt.Println(u, endereco)
	// fmt.Println(u.nome)
	// fmt.Println(u.idade)

	var u2 Usuario
	u2.nome = "Roberval"
	u.idade = 80

	fmt.Println(u2)
	// fmt.Println(u2.nome)
	// fmt.Println(u2.idade)

	u3 := Usuario{idade: 21}
	fmt.Println(u3)

}
