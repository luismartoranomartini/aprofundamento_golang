package main

import "fmt"

type Pilha[T any] struct {
	itens []T
}

func (p *Pilha[T]) Push(v T) {
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

func (p *Pilha[T]) Len() int { return len(p.itens) }

func main() {
	var p Pilha[int]
	p.Push(1)
	p.Push(2)
	p.Push(3)
	fmt.Println(p.Len(), p.itens)
	v, ok := p.Pop()
	fmt.Println(v, ok)
	v, ok = p.Pop()
	fmt.Println(v, ok)
	fmt.Println(p.Len(), p.itens)
	p.Pop()
	v, ok = p.Pop()
	fmt.Println(v, ok)
	fmt.Printf("%T\n", p)
}
