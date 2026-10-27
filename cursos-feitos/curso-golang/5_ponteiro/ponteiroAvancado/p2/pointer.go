package main

import "fmt"

func main() {
	var x *int
	a := 10
	*x = a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*x)
	fmt.Println(&x)
}
