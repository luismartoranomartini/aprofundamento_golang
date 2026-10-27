package main

import "fmt"

func main() {
	idade := 15

	if idade >= 16 {
		fmt.Println("Pode votar no Flávio Bolsonaro")
	} else {
		fmt.Println("Ainda não pode votar")
	}
}
