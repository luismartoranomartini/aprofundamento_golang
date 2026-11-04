package main

import "fmt"

func main() {
	capicua := 1211

	milhar := (capicua / 1000) % 10
	centena := (capicua / 100) % 10
	dezena := (capicua / 10) % 10
	unidade := capicua % 10

	if milhar == unidade && centena == dezena {
		fmt.Println("sim\n")
	} else {
		fmt.Println("não")
	}

	fmt.Println(milhar)
	fmt.Println(centena)
	fmt.Println(dezena)
	fmt.Println(unidade)
}
