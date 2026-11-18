package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	// Atribuição NÃO copia dados: b e a são dois descritores do MESMO array.
	b := a
	b[0] = 99
	fmt.Println(a, b) // [99 2 3] [99 2 3]

	// Clone de verdade, forma 1: make com o mesmo len + copy.
	c := make([]int, len(a))
	copy(c, a)
	c[1] = 77
	fmt.Println(a, c) // [99 2 3] [99 77 3]

	// Clone de verdade, forma 2: append sobre um slice nil, espalhando a origem com "...".
	d := append([]int(nil), a...)
	d[2] = 55
	fmt.Println(a, d) // [99 2 3] [99 2 55]

	// Compara o endereço do primeiro elemento: b compartilha com a; c e d não.
	fmt.Println(&a[0] == &b[0], &a[0] == &c[0], &a[0] == &d[0]) // true false false
}
