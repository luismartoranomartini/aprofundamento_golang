package main

import "fmt"

func main() {
	texto := "Estudo de Go em casa"
	vogais := 0

	for i := 0; i < len(texto); i++ {
		c := texto[i]

		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			vogais++
		}
	}
	fmt.Println(vogais)
}
