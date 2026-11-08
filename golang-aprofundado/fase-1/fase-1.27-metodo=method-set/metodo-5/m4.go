package main

import "fmt"

type Contador struct{ N int }

func (c *Contador) Nome() string {
	if c == nil {
		return "(nil)"
	}
	return "contador"
}

func (c *Contador) Valor() int { return c.N }

func main() {
	var p *Contador
	fmt.Println(p.Nome())
	defer func() {
		fmt.Println("recuperado:", recover())
	}()
	fmt.Println(p.Valor())
}
