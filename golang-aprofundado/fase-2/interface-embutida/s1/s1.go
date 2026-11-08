package main

import "fmt"

type Ponto struct{ X, Y int }

func (p Ponto) String() string { return fmt.Sprintf("(%d, %d)", p.X, p.Y) }

type Caixa struct{ N int }

func (c *Caixa) String() string { return fmt.Sprintf("Caixa#%d", c.N) }

func main() {
	p := Ponto{1, 2}
	fmt.Println(p)
	fmt.Printf("%v | %s | %+v\n", p, p, p)
	fmt.Println([]Ponto{{1, 2}, {3, 4}})

	c := Caixa{7}
	fmt.Println(c)  // valor: o String é de *Caixa, não entra no method set de Caixa
	fmt.Println(&c) // ponteiro: usa String
}
