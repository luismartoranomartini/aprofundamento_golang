package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  *string
	Idade int
}

func main() {
	var nome = "Pedro"

	pessoa1 := Pessoa{
		Nome:  &nome, // precisa estar em uma variável
		Idade: 26,
	}
	fmt.Println("---Forma Errada---")
	pessoa2 := pessoa1
	fmt.Println("Nome da pessoa 1:", *pessoa1.Nome)
	fmt.Println("Nome da pessoa 2:", *pessoa2.Nome)

	fmt.Println("Mudar o nome da pesso2")
	*pessoa2.Nome = "José"
	fmt.Println(*pessoa2.Nome)
	fmt.Println(*pessoa1.Nome) // compartilham o mesmo endereço de memória

	fmt.Println("---Forma Correta---")
	pessoa2.Nome = toPoiter("Davi")
	fmt.Println(*pessoa1.Nome)
	fmt.Println(*pessoa2.Nome)

	fmt.Println("---Deep Copy---")
	pessoa3 := Pessoa{
		Nome:  toPoiter("Leonardo"),
		Idade: 26,
	}

	pessoa4 := deepCopy(pessoa3)

	*pessoa4.Nome = "Lucas"
	fmt.Println("Nome da pessoa 3:", *pessoa3.Nome)
	fmt.Println("Nome da pessoa 3:", &pessoa3.Nome)
	fmt.Println("Nome da pessoa 4:", *pessoa4.Nome)
	fmt.Println("Nome da pessoa 4:", &pessoa4.Nome)

	// copy() // feita para suportar listas

	fmt.Println("---------------")
	listaDePessoas := []Pessoa{pessoa1, pessoa2, pessoa3, pessoa4}
	listaDePessoas2 := deepCopyList(listaDePessoas)

	fmt.Println("Lista 1: ", listaDePessoas)
	fmt.Println("Lista 2: ", listaDePessoas2)

}
func toPoiter(s string) *string {
	return &s
}

func deepCopy(origem Pessoa) Pessoa {
	var destino Pessoa
	destino.Idade = origem.Idade
	destino.Nome = toPoiter(*origem.Nome)

	return destino
}

func deepCopyList(origem []Pessoa) []Pessoa {
	var destino = make([]Pessoa, len(origem))

	for i, pessoa := range origem {
		destino[i] = deepCopy(pessoa)
	}
	return destino
}
