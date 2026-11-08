package main

import "fmt"

func contador() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

func main() {
	a := contador()
	b := contador()
	fmt.Println(a(), a(), a())
	fmt.Println(b())
	fmt.Println(a())
}
