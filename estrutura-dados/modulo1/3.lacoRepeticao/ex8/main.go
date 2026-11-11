package main

import "fmt"

func main() {
	s := "olá"
	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte: %d\n", s[i])
	}

	for _, r := range s {
		fmt.Printf("Rune: %d\n", r)
	}
}
