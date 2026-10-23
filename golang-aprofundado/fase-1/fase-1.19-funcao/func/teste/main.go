package main

import "fmt"

func f(xs ...int) (n int, s int) {
	for _, x := range xs {
		n++
		s += x
	}
	return
}

func main() {
	fmt.Println(f(4, 5, 6))
	fmt.Println(f())
}
