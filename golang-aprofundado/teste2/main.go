package main

import "fmt"

func main() {
	var valor any = 42
	numero := valor.(int)
	texto := valor.(string)

	fmt.Printf("Numero:%T\n", numero)
	fmt.Printf("Texto: %T\n", texto)
}
