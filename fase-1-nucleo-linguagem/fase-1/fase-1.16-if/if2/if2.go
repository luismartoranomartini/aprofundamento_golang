package main

import "fmt"

func main() {
	x := 1
	if true {
		x := 2
		fmt.Println("dentro:", x)
	}
	fmt.Println("fora:", x)
}
