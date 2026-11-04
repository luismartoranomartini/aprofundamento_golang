package main

import "fmt"

func main() {
	numero := 55

	dezena := (numero / 10)
	unidade := numero % 10

	if dezena == unidade {
		fmt.Println("são iguais")
	} else {
		fmt.Println("são diferentes")
	}
}
