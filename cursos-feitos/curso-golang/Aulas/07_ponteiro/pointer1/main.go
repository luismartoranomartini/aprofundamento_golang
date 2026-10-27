package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x

	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)
}
