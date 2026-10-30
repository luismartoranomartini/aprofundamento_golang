package main

import "fmt"

func main() {
	numero := 4444

	milhar := numero / 1000
	centena := (numero / 100) % 100
	dezena := (numero / 10) % 10
	unidade := numero % 10

	fmt.Printf("Número:  %d\n", numero)
	fmt.Printf("Milhar:  %d\n", milhar)
	fmt.Printf("Centena: %d\n", centena)
	fmt.Printf("Dezena:  %d\n", dezena)
	fmt.Printf("Unidade: %d\n", unidade)
}
