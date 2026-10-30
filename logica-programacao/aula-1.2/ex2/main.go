package main

import "fmt"

func main() {
	n := 5

	if n%2 == 0 {
		fmt.Printf("%d é par\n", n)
	} else {
		fmt.Printf("%d é ímpar\n", n)
	}
}
