package main

import "fmt"

func main() {
	notas := map[string]int{"Ana": 9, "Caop": 0}

	// Sem o segundo valor, "ausente" e "zero" parecem iguais.
	fmt.Println(notas["Caio"], notas["Zé"]) // 0 0

	// Com o segundo valor (ok), dá para separar os dois casos.
	v, ok := notas["Caio"]
	fmt.Println(v, ok)

	v, ok = notas["Zé"]
	fmt.Println(v, ok)

	// Forma comum: declarar e testar na mesma linha do if.
	if nota, achou := notas["Ana"]; achou {
		fmt.Println("Ana tem nota", nota) // Ana tem 9
	}

}
