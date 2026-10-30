package main

import "fmt"

func main() {
	n := 330

	if (n / 100 % 10) == (n % 10) {
		fmt.Println("É palíndromo")
	} else {
		fmt.Println("Não é palíndromo")
	}
}
