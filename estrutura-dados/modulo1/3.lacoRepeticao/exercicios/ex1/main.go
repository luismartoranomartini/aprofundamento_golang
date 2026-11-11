package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Crescente: ", i)
	}

	for i := 10; i >= 1; i-- {
		fmt.Println("Decrescente: ", i)
	}

	for i := 2; i <= 10; i += 2 {
		fmt.Printf("Pares: %d\n", i)
	}
}
