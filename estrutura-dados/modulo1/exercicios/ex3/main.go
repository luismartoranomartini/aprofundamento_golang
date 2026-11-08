package main

import "fmt"

func main() {
	palavra := "coração"

	fmt.Println("len:", len(palavra))

	for i := 0; i < len(palavra); i++ {
		fmt.Printf("byte em hexadecimal: %x\n", palavra[i])

	}
}
