package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println("defer direto, x =", x)
	defer func() { fmt.Println("defer clousure, x =", x) }()
	x = 2
	fmt.Println("corpo, x =", x)
}
