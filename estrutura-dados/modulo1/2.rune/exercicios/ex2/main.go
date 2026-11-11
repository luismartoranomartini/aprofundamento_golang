package main

import "fmt"

func main() {
	s := "coração"

	for i, r := range s {
		fmt.Printf("%d: %d (%c)\n", i, r, r)
	}
}
