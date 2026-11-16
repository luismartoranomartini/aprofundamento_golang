package main

import "fmt"

func main() {
	texto := "Estudo de Go em casa"
	vogais := 0

	for _, c := range texto {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vogais++
		}
	}
	fmt.Println(vogais)
}
