package main

import "fmt"

func classificarIdade(idade int) string {
	if idade < 0 {
		return "inválida"
	} else if idade < 12 {
		return "criança"
	} else if idade < 18 {
		return "adolescente"
	} else {
		return "adulto"
	}
}

func main() {
	fmt.Println(classificarIdade(-1))
	fmt.Println(classificarIdade(5))
	fmt.Println(classificarIdade(12))
	fmt.Println(classificarIdade(17))
	fmt.Println(classificarIdade(18))
}
