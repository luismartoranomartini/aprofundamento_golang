package main

import "fmt"

func main() {
laco:
	for i := 0; i < 5; i++ {
		switch i {
		case 2:
			break laco
		default:
			fmt.Println("i =", i)
		}
	}
	fmt.Println("fim")
}
