package main

import "fmt"

func main() {
	for i, r := range "olá" {
		fmt.Printf("i=%d r=%d %c\n", i, r, r)
	}
}
