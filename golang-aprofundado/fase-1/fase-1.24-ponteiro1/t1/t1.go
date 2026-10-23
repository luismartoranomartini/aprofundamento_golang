package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Printf("tipo de x: %T\n", x)
	fmt.Printf("tipo de p: %T\n", p)

	fmt.Println("x =", x)
	fmt.Println("*p =", *p)

	*p = 20
	fmt.Println("x depois de *p = 20:", x)

	x = 30
	fmt.Println("*p depois de x = 30:", *p)

	fmt.Println("p == &x", p == &x)
}
