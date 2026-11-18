package main

import "fmt"

func main() {
	// Antes de rodar, escreva aqui, em comentário, as duas linhas de saída
	// e o motivo de cada valor (conteúdo, len e cap de s e de t).

	s := []int{10, 20, 30, 40, 50}
	t := append(s[:1], s[2:]...)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
}
