package main

import "fmt"

func main() {
	texto := "palavra"
	tamanho := len(texto)
	fmt.Println("Quantidade: ", len(texto))

	i := 0

	for i < tamanho {


		if string(texto[i]) == "r"{
			i++
			break
		}

		fmt.Println(string(texto[i]))
		i++

	}
}
