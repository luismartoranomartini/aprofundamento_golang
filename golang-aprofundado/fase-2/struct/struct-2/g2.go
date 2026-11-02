package main

import "fmt"

type Ponto struct{ X, Y int }

func main() {
	a := Ponto{1, 2}
	b := Ponto{Y: 2, X: 1}
	c := &a
	d := &Ponto{1, 2}
	fmt.Println(a == b, c == d, *c == *d)

	m := map[Ponto]string{{1, 2}: "um"}
	m[Ponto{1, 2}] = "dois"
	fmt.Println(len(m), m[a])
}
