package main

import "fmt"

func main() {
	for i, r := range "coração" {
		fmt.Printf("%d:%c\n", i, r)
	}
}
