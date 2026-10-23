package main

import "fmt"

type Contador struct{ N int }

func (c *Contador) Inc() { c.N++ }

type Painel struct{ Contador }   // embute T
type PainelP struct{ *Contador } // embute *T

type Incrementador interface{ Inc() }

func main() {
	var p Painel
	p.Inc()                  // p é endereçável: (&p).Inc()
	fmt.Println(p.N)         // 1
	var a Incrementador = &p // *Painel tem Inc
	a.Inc()
	fmt.Println(p.N) // 2
	pp := PainelP{&Contador{}}
	var b Incrementador = pp // PainelP (valor) também tem Inc
	b.Inc()
	fmt.Println(pp.N)
	var c Incrementador = p
	fmt.Println(c)
}
