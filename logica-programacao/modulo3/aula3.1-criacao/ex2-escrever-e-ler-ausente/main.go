package main

import "fmt"

func main() {
	notas := map[string]int{"Ana": 8}

	notas["Bia"] = 10              // chave nova: acrescenta o par
	notas["Ana"] = 9               // chave que já existe: substitui o valor
	fmt.Println(notas, len(notas)) // map[Ana:9 Bia:10] 2

	fmt.Println(notas, notas["Ana"]) // 9
	fmt.Println(notas["Zé"])         // 0: chave ausente devolve o valor zero do tipo
	fmt.Println(len(notas))          // 2: ler uma chave ausente NÃO cria a chave
}
