package main

import "fmt"

type Pessoa struct {
	Nome  *string
	Idade int
}

func main() {
	var nome = "Pedro"
	pessoa := Pessoa{
		Nome:  &nome,
		Idade: 20,
	}
	// errada
	pessoa2 := pessoa
	fmt.Println("Nome da pessoa 1:", *pessoa.Nome)
	fmt.Println("Nome da pessoa 2:", *pessoa2.Nome)
	*pessoa2.Nome = "Davi"
	fmt.Println(*pessoa.Nome)
	fmt.Println(*pessoa2.Nome)

	//Forma correta
	fmt.Println("---Com ponteiros---")
	pessoa2.Nome = toPointer("Jair")
	fmt.Println(*pessoa.Nome)
	fmt.Println(*pessoa2.Nome)

	fmt.Println("---Com função---")
	pessoa3 := Pessoa{
		Nome:  toPointer("Kiko"),
		Idade: 26,
	}

	// deep Copy
	pessoa4 := deepCopy(pessoa3)
	fmt.Println("Nome da pessoa 3:", *pessoa3.Nome)
	fmt.Println("Nome da pessoa 4:", *pessoa4.Nome)
	*pessoa4.Nome = "Sebastião"
	fmt.Println("Nome da pessoa 3 modificado:", *pessoa3.Nome)
	fmt.Println("Nome da pessoa 4 modificado:", *pessoa4.Nome)

	fmt.Println("------------------------------")
	listaDePessoas1 := []Pessoa{pessoa, pessoa2, pessoa3, pessoa4}
	listaDePessoas2 := DeepCopyList(listaDePessoas1)
	fmt.Println("Lista 1", listaDePessoas1)
	fmt.Println("Lista 2", listaDePessoas2)
}

func toPointer(s string) *string {
	return &s
}

func deepCopy(origem Pessoa) Pessoa {
	var destino Pessoa
	destino.Idade = origem.Idade
	destino.Nome = toPointer(*origem.Nome)
	return destino
}

func DeepCopyList(origem []Pessoa) []Pessoa {
	var destino = make([]Pessoa, len(origem))
	for i, pessoa := range origem {
		destino[i] = deepCopy(pessoa)
	}
	return destino
}
