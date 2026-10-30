package main

import "fmt"

func main() {
	var peso float64 = 72.5
	var altura float64 = 1.75
	imc := peso / (altura * altura)

	if imc < 18.5 {
		fmt.Println("Abaixo do peso")
	} else if imc < 25.0 {
		fmt.Println("Peso normal")
	} else if imc < 30.0 {
		fmt.Println("sobrepeso")
	} else {
		fmt.Println("Obesidade")
	}
}
