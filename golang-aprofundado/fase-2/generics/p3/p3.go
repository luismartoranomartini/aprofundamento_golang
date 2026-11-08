package main

import "fmt"

type Pilha[T any] struct{ itens []T }

func (p *Pilha[T]) Push(v T) { p.itens = append(p.itens, v) }
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
	var p Pilha[string]
	p.Push("a")
	p.Push("b")
	p.Push("c")
	x, _ := p.Pop()
	p.Push("d")
	fmt.Println(x, len(p.itens))
	for {
		v, ok := p.Pop()
		if !ok {
			fmt.Printf("%q %v\n", v, ok)
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Printf("%T\n", p)
}
