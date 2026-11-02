package main

import "fmt"

type Par[K comparable, V any] struct {
	Chave K
	Valor V
}

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

func main() {
	s := &Pilha[string]{}
	s.Push("a")
	s.Push("b")
	x, _ := s.Pop()
	fmt.Println(x)
	fmt.Printf("%T %T\n", s, *s)

	pr := Par[string, int]{"idade", 30}
	fmt.Printf("%v %T\n", pr, pr)

	var pp Pilha[Par[string, int]] // pilha de pares
	pp.Push(pr)
	t, ok := pp.Pop()
	fmt.Println(t, ok)
}
