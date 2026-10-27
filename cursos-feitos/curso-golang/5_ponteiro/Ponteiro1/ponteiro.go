package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 10
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro) //ponteiro é o endereço da variável3

	variavel3 = 150

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}
