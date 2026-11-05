package main

import "fmt"

func entrar(idade int, temIngresso bool) string {
	if idade < 0 {
		return "idade inválida"
	}
	if !temIngresso {
		return "sem ingresso"
	}

	if idade < 18 {
		return "menor de idade"
	}
	return "pode entrar"
}

func main() {
	fmt.Println(entrar(-1, true))
	fmt.Println(entrar(20, false))
	fmt.Println(entrar(10, true))
	fmt.Println(entrar(20, true))
}
