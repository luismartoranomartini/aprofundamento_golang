package main

import "fmt"

func main() {
	fmt.Println("início")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("fim")
}
