package main

import (
	model "08_struct/model"
	"fmt"
	"time"
)

func main() {

	fmt.Println("Iniciando...")

	endereco := model.Endereco{
		Rua:    "Rua X",
		Numero: 15,
		Cidade: "Bauru",
	}
	pessoa := model.Pessoa{
		Nome:           "Luís Martini",
		Endereco:       endereco,
		DataNascimento: time.Date(1977, 07, 14, 0, 0, 0, 0, time.Local),
	}

	autoMoto := model.Auto{
		Ano:    2026,
		Placa:  "XPTO",
		Modelo: "GC",
	}

	moto := model.Moto{
		Auto:        autoMoto,
		Cilindradas: 125,
	}

	fmt.Println(moto)
	fmt.Println(moto.Ano)
	// idade := model.CalculaIdade(pessoa)

	fmt.Println(pessoa)
	fmt.Println(endereco)
	// pessoa.IdadeAtual() // elegante
	pessoa.ProcessaIdade()
	fmt.Println(pessoa.Idade)

}
