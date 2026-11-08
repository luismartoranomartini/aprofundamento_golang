package main

import "fmt"

type Contador struct{ N int }

func (c *Contador) Inc()    { c.N++ }
func (c Contador) Ver() int { return c.N }

type Incrementador interface{ Inc() }
type Leitor interface{ Ver() int }

func main() {
	c := Contador{}
	var a Incrementador = &c
	var b Leitor = &c
	var d Leitor = c
	a.Inc()
	a.Inc()
	fmt.Println(b.Ver(), d.Ver(), c.N)
}
