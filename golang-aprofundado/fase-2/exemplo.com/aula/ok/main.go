package main

import (
	"fmt"

	"exemplo.com/aula/conta"
)

func main() {
	c := conta.Nova("Ana")
	c.Depositar(100)
	c.Depositar(-5)
	conta.Dobrar(c)
	fmt.Println(c.Titular, c.Saldo())
	fmt.Println(conta.Resumo(c))
	fmt.Printf("%T\n", c)
}
