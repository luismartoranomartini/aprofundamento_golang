package main

import "fmt"

func main() {

	// var num float32 = 999.00
	var num float32 = 5000.00

	if num < 1000.00 {
		num -= num * 0.08
		fmt.Println("Menor que 1000 = ", num)
	} else if num >= 1000.00 && num < 5000.00 {
		num -= num * 0.15
		fmt.Println("Entre 1000 e 4999 = ", num)
	} else {
		num -= num * 0.10
		fmt.Println("Maior ou igual a 5000 (Desconto de 10%) = ", num)
	}

}
