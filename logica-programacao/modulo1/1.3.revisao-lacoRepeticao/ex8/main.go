package main

import "fmt"

func main() {
	a, b := 0, 1
	for range 8 {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}
