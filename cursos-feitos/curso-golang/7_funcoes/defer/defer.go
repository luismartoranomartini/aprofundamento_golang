package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado.")
	fmt.Println("Entrando para verificar")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {

	fmt.Println(alunoAprovado(8, 7))
	// fmt.Println(alunoAprovado(5, 3))
}
