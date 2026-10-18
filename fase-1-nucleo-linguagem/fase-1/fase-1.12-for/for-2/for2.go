package main

import "fmt"

func main() {
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// o tipo de i segue o tipo do inteiro
	for i := range uint8(3) {
		fmt.Printf("%T ", i)
	}
	fmt.Println()

	// range de inteiro avalia n UMA vez
	n := 3
	vezes := 0
	for range n {
		n = 10
		vezes++
	}
	fmt.Println(vezes, n)
}
