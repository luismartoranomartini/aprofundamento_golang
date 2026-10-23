package main

import "fmt"

type Contador struct {
	N int
}

func (c Contador) IncV()    { c.N++ }
func (c *Contador) inc()    { c.N++ }
func (c Contador) Ver() int { return c.N }

func main() {
	c := Contador{}
	c.IncV()
	fmt.Println("depois de IncV:", c.N)
	c.inc()
	c.inc()
	fmt.Println("depois de Inc x2:", c.N)

	p := &c
	fmt.Println("p.Ver():", p.Ver())
	p.inc()
	fmt.Println("c.Ver():", c.Ver())
}
