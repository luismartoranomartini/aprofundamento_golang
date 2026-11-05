package main

import "fmt"

func par(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(par(2))
	fmt.Println(par(3))
	fmt.Println(par(4))
	fmt.Println(par(5))
}
