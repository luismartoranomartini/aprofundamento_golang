package main

import "fmt"

func main() {

	var hoje int
	var passa int

	fmt.Print("Digite um dia da semana de 0 a 6: ")
	fmt.Scan(&hoje)

	fmt.Print("Digite quantos dias se passaram: ")
	fmt.Scan(&passa)

	dia := (hoje + passa) % 7

	if dia == 0 || dia == 6 {
		fmt.Println("É final de semana")
	} else {
		fmt.Println("Não é final de semana")
	}
}
