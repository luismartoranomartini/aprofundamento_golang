package main

import "fmt"

func main() {
	var valor any = "Olá, Go!"

	texto, ok := valor.(string)
	if ok {
		fmt.Println("Sucesso: ", texto)
	}

	numero, ok := valor.(int)
	fmt.Println("numero:", numero, "| Sucesso:", ok)
}
