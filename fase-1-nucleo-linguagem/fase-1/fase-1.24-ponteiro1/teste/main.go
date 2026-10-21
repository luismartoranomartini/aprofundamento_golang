package main

import "fmt"

func main() {
	x := 1
	y := 2
	p := &x
	q := &y
	*p = *q
	q = p
	*q = 10
	fmt.Println(x, y)
	fmt.Println(*p, *q)
	fmt.Println(p == q)
}
