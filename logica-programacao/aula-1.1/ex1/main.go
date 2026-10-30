package main

import "fmt"

func main() {
	var n int
	var f float64
	var s string
	var b bool

	fmt.Printf("%T %v\n", n, n)
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %q\n", s, s)
	fmt.Printf("%T %v\n", b, b)

}
