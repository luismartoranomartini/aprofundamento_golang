package main

import "fmt"

func main() {
	var palavra string
	var inicio int
	var final int

	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&palavra)
	fmt.Println("Digite o início da divisão")
	fmt.Scan(&inicio)
	fmt.Println("Digite o final da divisão")
	fmt.Scan(&final)

	if inicio == 0 && final == 0 {
		fmt.Println(palavra)
	} else {
		fmt.Println(palavra[inicio:(final - 1)])
	}

}
