package main

import "fmt"

func main() {

	texto := "palavra"
	fmt.Println("Quantidade: ", len(texto))

	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			// break
			continue
			
		}
		fmt.Println(string(texto[i]))

	}

}
