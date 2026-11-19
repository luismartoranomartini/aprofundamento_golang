package main

import (
	"fmt"
	"strconv"
)

func main() {
	var valor any = 42

	numero := valor.(int)
	texto := strconv.Itoa(numero)

	fmt.Printf("Numero: %T (%d)\n", numero, numero)
	fmt.Printf("Texto: %T (%q)\n", texto, texto)
}
