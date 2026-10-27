package main

import "fmt"

func main() {

	resultado := Soma(1, 2)
	fmt.Println(resultado)


	
}

func Soma(n1 int, n2 int) int {
	resultado := n1 + n2
	return resultado
}