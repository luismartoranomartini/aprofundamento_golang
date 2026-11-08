package main

import (
	"fmt"

	"exemplo.com/aula/conta"
)

func main() {
	c := conta.Nova("Bia")
	c.Depositar(30)
	c.Depositar(0)
	c.Depositar(-1)
	conta.Dobrar(c)
	conta.Dobrar(c)
	fmt.Println(c.Saldo(), conta.Resumo(c))
	fmt.Printf("%+v\n", *c)
	fmt.Printf("%T\n", *c)
}
