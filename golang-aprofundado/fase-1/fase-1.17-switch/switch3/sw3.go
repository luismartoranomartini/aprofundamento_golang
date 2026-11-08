package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		switch i {
		case 1:
			break // sai do switch, NÃO do for
		default:
			fmt.Println("i =", i)
		}
	}
	fmt.Println("fim")
}
