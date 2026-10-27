package main

import "fmt"

func main() {
	var imc float64

	imc = 18

	if imc >= 30 {
		fmt.Println("Obesidade")
	} else if imc >= 25 {
		fmt.Println("Sobrepeso")
	} else if imc >= 18.5 {
		fmt.Println("peso normal")
	} else {
		fmt.Println("abaixo do peso")
	}
}
