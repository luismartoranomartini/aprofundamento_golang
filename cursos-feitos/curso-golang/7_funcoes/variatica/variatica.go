package main

import "fmt"

func soma(numeros ...int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println(total)
}

func escrever(txt string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(txt, numero)
	}
}

func main() {
	soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	soma(1, 2, 3, 4, 5)
	soma()

	escrever("Luís Martini", 1, 2, 3, 4, 5)
}
