package main

import "fmt"

func main() {
	// 1. Literal: já nasce com pares chave: valor.
	notas := map[string]int{"Ana": 8, "Bia": 10}

	// 2. make: nasce vazio, mas pronto para receber pares.
	idades := make(map[string]int)

	// 3. var sem valor: nasce nil (não existe tabela por trás).
	var vazio map[string]int

	fmt.Println(notas, len(notas))
	fmt.Println(idades, len(idades))
	fmt.Println(vazio, len(vazio))
	fmt.Println(idades == nil, vazio == nil)

}
