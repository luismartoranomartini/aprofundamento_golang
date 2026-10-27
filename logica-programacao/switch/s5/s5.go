package main

import "fmt"

func main() {
	var ano int

	fmt.Print("Digite o ano: ")
	_, err := fmt.Scan(&ano)
	if err != nil {
		fmt.Println("Ano inválido.")
		return
	}

	switch {
	case ano%400 == 0: // 1) O mais importante
		fmt.Println("Bissexto")
	case ano%100 == 0: // 2) multiplo de 100, mas não de 400
		fmt.Println("Nao bissexto")
	case ano%4 == 0: // 3) múltiplo de 4, mas não de 100
		fmt.Println("Bissexto")
	default: // 4) o resto
		fmt.Println("Nao bissexto")
	}
}
