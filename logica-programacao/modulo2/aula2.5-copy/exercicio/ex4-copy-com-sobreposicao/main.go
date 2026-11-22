package main

import "fmt"

func main() {
	// Antes de rodar, escreva aqui, em comentário, a saída esperada e o
	// motivo do valor de n e do conteúdo final de cada slice.
	a := []int{1, 2, 3, 4, 5}
	n1 := copy(a[2:], a)
	fmt.Println(n1, a)

	b := []int{1, 2, 3, 4, 5}
	n2 := copy(b, b[3:])
	fmt.Println(n2, b)
}
