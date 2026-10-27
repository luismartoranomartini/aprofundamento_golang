package main

import (
	"fmt"
)


// var numeroMagico = 5

func main() {

	soma, sub, div, mult := Operacao(1, 2)
	fmt.Println(soma)
	fmt.Println(sub)
	fmt.Println(div)
	fmt.Println(mult)

}

func Operacao(n1 int, n2 int) (
	soma, 
	sub, 
	div, 
	mult int) {


	soma = n1 + n2 
	sub = n1 - n2
	div = n1 / n2
	mult = n1 * n2

	return
}

func init() {
	fmt.Println("Função Init")
}

