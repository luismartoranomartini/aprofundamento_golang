package main

import "fmt"

type Somador interface{ Soma() int }

type A struct{ N int }

func (a A) Soma() int { return a.N }

type B struct{ N int }

func (b *B) Soma() int { b.N++; return b.N }

func total(ss []Somador) int {
	t := 0
	for _, s := range ss {
		t += s.Soma()
	}
	return t
}

func main() {
	a := A{10}
	b := B{20}
	ss := []Somador{a, &b}
	a.N = 99
	fmt.Println(total(ss), total(ss))
	fmt.Println(a.N, b.N)
}
