package main

import "fmt"

func ajuda() {
	fmt.Println("ajuda recover:", recover())
}

func main() {
	defer func() {
		fmt.Println("closure recover:", recover())
	}()
	defer func() {
		ajuda()
	}()
	panic("x")
}
