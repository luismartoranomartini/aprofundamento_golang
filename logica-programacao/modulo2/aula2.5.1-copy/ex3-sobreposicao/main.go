package main

import "fmt"

func main() {
	// copy funciona mesmo quando origem e destino usam o mesmo array.
	s := []int{1, 2, 3, 4, 5}
	copy(s[1:], s) // desloca tudo uma posição à direita; o último (5) se perde
	fmt.Println(s) // [1 1 2 3 4]

	t := []int{1, 2, 3, 4, 5}
	copy(t, t[1:]) // desloca tudo uma posição à esquerda; o último fica repetido
	fmt.Println(t) // [2 3 4 5 5]
}
