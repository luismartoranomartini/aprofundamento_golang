package main

import "fmt"

func main() {
	// Reserva 4 posições de uma vez: len 0, cap 4.
	r := make([]int, 0, 4)

	// Guarda um descritor do começo do array: base[0] é o primeiro elemento.
	r = append(r, 10)
	base := r

	for i := 2; i <= 4; i++ {
		r = append(r, i*10)
	}
	// len 4 e cap 4: o cap é o mesmo da reserva. Nenhum append passou do cap,
	// então o array nunca trocou.
	fmt.Println(r, len(r), cap(r)) // [10 20 30 40] 4 4
	fmt.Println(&r[0] == &base[0]) // true

	// Sem reserva: se o cap muda durante os appends, houve array novo.
	var x []int
	mudou := false
	ultimo := cap(x)
	for i := 1; i <= 4; i++ {
		x = append(x, i*10)
		if cap(x) != ultimo {
			mudou = true
			ultimo = cap(x)
		}
	}
	fmt.Println(x, len(x), mudou) // [10 20 30 40] 4 true
}
