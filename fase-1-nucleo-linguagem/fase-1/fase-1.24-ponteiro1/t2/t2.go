package main

import "fmt"

func main() {
	a := 1
	p := &a
	q := p
	*q = 99
	fmt.Println("a =", a, "*p =", *p, "*q =", *q)
	fmt.Println("p == q", p == q)

	b := 99
	r := &b
	fmt.Println("p == r:", p == r, "| *p == *r:", *p == *r)

	pp := &p
	fmt.Printf("tipo de pp: %T\n", pp)
	**pp = 7
	fmt.Println("a =", a)
}
