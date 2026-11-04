package main

import "fmt"

func main() {
	numero := 359
	centena := (numero / 100) % 10
	dezena := (numero / 10) % 10
	unidade := numero % 10

	if centena < dezena && dezena < unidade {
		fmt.Println("sim")
	} else {
		fmt.Println("não")
	}
}
