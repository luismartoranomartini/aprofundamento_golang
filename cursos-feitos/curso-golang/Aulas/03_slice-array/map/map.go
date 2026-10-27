package main

import (
	"fmt"
)

func main() {

	// m := map[string]int {"SP": 1000, "CG": 900}

	m1 := make(map[string]int)
	m1["SP"] = 10000000
	m1["CG"] = 9000000

	valor, existe := m1["RJ"]

	if existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Chave não existe")
	}

	if valor, existe := m1["RJ"]; existe {
		fmt.Println(valor, existe)
	} else {
		fmt.Println("Chave não existe")
	}

	// fmt.Println(m1)
	// fmt.Println(valor, existe)
}
