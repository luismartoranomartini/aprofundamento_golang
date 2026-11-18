package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	b := a[2:5]                    // forma simples: cap vai até o fim do array
	fmt.Println(b, len(b), cap(b)) // [2 3 4] 3 8

	c := a[2:5:5]                  // max = 5: len = 5-2 = 3, cap = 5-2 = 3
	fmt.Println(c, len(c), cap(c)) // [2 3 4] 3 3

	d := a[2:5:7]                  // max = 7: len = 3, cap = 7-2 = 5
	fmt.Println(d, len(d), cap(d)) // [2 3 4] 3 5

	e := a[:0:2]                   // low 0, high 0, max 2: len 0, cap 2
	fmt.Println(e, len(e), cap(e)) // [] 0 2

	f := a[0:0:0]                  // len 0, cap 0: nenhuma posição pode ser usada por append
	fmt.Println(f, len(f), cap(f)) // [] 0 0

	// Regra de validade: 0 <= low <= high <= max <= cap(a). Fora dela dá erro.
	// Exemplo inválido (não executar): a[2:5:4] tem high > max.
}
