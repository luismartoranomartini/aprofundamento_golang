package main

import "fmt"

func main() {
	n := "Gopher"

	for i := len(n) - 1; i >= 0; i-- {
		fmt.Printf("%d: %s\n", i, string(n[i]))
	}
}
