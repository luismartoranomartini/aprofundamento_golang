package main

import "fmt"

func main() {
	var n int
	n = 12
	for i := 0; i < n; i += 2 {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Printf(" %d ", i)
	}
	fmt.Println()
}
