package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func novoContador() *int {
	c := 100
	return &c
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

	p := novoContador()
	q := novoContador()
	*p += 1
	fmt.Println(*p, *q, p == q)
}
