package main

import "fmt"

func main() {
	nota := 10

	if nota >= 7 {
		fmt.Println("aprovado")
	} else if nota >= 5 {
		fmt.Println("recuperação")
	} else {
		fmt.Println("reprovado")
	}
}
