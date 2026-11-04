package main

import "fmt"

func idadeValida(idade int) string {
	if idade > 0 {
		return "Idade válida"
	} else {
		return "idade inválida"
	}
}

func main() {
	fmt.Println(idadeValida(10))

	fmt.Println(idadeValida(-10))
}
