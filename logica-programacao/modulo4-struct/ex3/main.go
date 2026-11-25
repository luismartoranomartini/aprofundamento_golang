package main

import "fmt"

type Conta struct {
	Titula string
	Saldo  float64
}

func depositar(c Conta, valor float64) Conta {
	c.Saldo += valor
	return c
}

func main() {
	c1 := Conta{"Ana", 100}

	depositar(c1, 50)
	fmt.Println("c1 depois:", c1)

	c2 := depositar(c1, 50)
	fmt.Println("c1:", c1, "| c2:", c2)

	d := c1
	d.Saldo = 0
	fmt.Println("c1:", c1, "| d:", d)
}
