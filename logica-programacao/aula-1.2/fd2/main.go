package main

import "fmt"

func main() {
	numero := 55

	dezena := numero / 10
	unidade := numero % 10

	par := dezena + unidade

	if par%2 == 0 {
		fmt.Println("É par:", par)
	} else {
		fmt.Println("É impar")
	}

}
