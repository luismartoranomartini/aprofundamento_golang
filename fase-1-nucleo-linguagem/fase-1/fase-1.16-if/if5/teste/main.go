package main

import "fmt"

func main() {
	x := 10
	if x := 20; x > 15 {
		x++
		fmt.Println("a", x)
	}
	fmt.Println("b", x)
}
