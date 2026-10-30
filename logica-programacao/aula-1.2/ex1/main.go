package main

import "fmt"

func main() {
	n := 0
	if n > 0 {
		fmt.Println("positivo")
	} else if n < 0 {
		fmt.Println("negativo")
	} else {
		fmt.Println("zero")
	}
}
