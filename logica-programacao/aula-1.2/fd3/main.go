package main

import "fmt"

func main() {
	numero := 7

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("ambos")
	} else if numero%5 == 0 {
		fmt.Println("só 5")
	} else if numero%3 == 0 {
		fmt.Println("só 3")
	} else {
		fmt.Println("nenhum")
	}
}
