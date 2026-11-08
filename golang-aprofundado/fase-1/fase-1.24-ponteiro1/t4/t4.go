package main

import "fmt"

func main() {
	var p *int
	x := 5
	if p == nil {
		p = &x
	}
	fmt.Println(*p)

	n := new(int)
	fmt.Println(*n)
	*n = 42
	fmt.Println(*n)

}
