package main

import "fmt"

func main() {
	var p *int
	fmt.Println("p == nil", p == nil)
	fmt.Println("*p:", p)

	defer func() {
		fmt.Println("recuperado:", recover())
	}()
	fmt.Println(*p)
}
