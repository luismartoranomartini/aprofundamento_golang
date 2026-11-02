package main

import (
	"fmt"
	"strconv"
)

type Pilha[T any] struct {
	itens []T
}

type PilhaInt struct {
	itens []int
}

type PilhaAny struct {
	itens []any
}

func (p *Pilha[T]) Push(v T) {
	p.itens = append(p.itens, v)
}

func (p *PilhaInt) Push(v int) {
	p.itens = append(p.itens, v)
}

func (p *PilhaAny) Push(v any) {
	p.itens = append(p.itens, v)
}

func (p *Pilha[T]) Pop() (T, bool) {
	var zero T
	if len(p.itens) == 0 {
		return zero, false
	}

	i := len(p.itens) - 1
	v := p.itens[i]
	p.itens = p.itens[:i]
	return v, true
}

func (p *PilhaAny) Pop() (any, bool) {
	var zero any
	if len(p.itens) == 0 {
		return zero, false
	}

	i := len(p.itens) - 1
	v := p.itens[i]
	p.itens = p.itens[:i]
	return v, true
}

func (p *PilhaInt) Pop() (int, bool) {
	var zero int
	if len(p.itens) == 0 {
		return zero, false
	}

	i := len(p.itens) - 1
	v := p.itens[i]
	p.itens = p.itens[:i]
	return v, true
}

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, 0, len(s))
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func (p *Pilha[T]) Len() int {
	return len(p.itens)
}

func (p *PilhaInt) Len() int {
	return len(p.itens)
}

func (p *PilhaAny) Len() any {
	return len(p.itens)
}

func main() {
	var g Pilha[string]
	g.Push("a")
	g.Push("b")
	v, ok := g.Pop()
	fmt.Println(v, ok, g.Len())

	r := Map([]int{1, 2, 3}, func(n int) string { return strconv.Itoa(n * n) })
	fmt.Printf("%v %T\n", r, r)

	var pi PilhaInt
	pi.Push(1)
	pi.Push(2)
	n, _ := pi.Pop()
	fmt.Println(n+1, pi.Len())

	var pa PilhaAny
	pa.Push(1)
	pa.Push("dois")
	x, _ := pa.Pop()
	fmt.Printf("%v %T\n", x, x)
	y, _ := pa.Pop()
	fmt.Println(y.(int) + 1)

	func() {
		defer func() { fmt.Println("recuperado:", recover()) }()
		pa.Push("texto")
		z, _ := pa.Pop()
		fmt.Println(z.(int) + 1)
	}()
}
