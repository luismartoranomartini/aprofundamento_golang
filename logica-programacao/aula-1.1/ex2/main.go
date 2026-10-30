package main

import "fmt"

func main() {
	a, b := 10, 20

	temp := a
	a = b
	b = temp

	fmt.Printf("Depois (a): a = %d, b = %d\n\n", a, b)

	a, b = b, a
	fmt.Println(a, b)
}
