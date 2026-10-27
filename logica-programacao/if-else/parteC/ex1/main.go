package main

import "fmt"

func main() {
	var hora int

	hora = 23

	if hora >= 0 && hora <= 11 {
		fmt.Println("Bom dia!")
	} else if hora >= 12 && hora <= 17 {
		fmt.Println("Boa tarde!")
	} else {
		fmt.Println("Boa noite")
	}

}
