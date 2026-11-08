package main

import "fmt"

func main() {
	a, err := 1, error(nil)
	b, err := 2, fmt.Errorf("x")
	fmt.Println(a, b, err)
}
