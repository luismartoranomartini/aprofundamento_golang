package main

import "fmt"

func main() {
	s := "aé世😀"
	for _, r := range s {
		fmt.Printf("Caracter %c | Unicode U+%04X | quantidade %d\n", r, r, len(string(r)))
	}
}
