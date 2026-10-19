package main

import "fmt"

func main() {
	i := 0
inicio:
	if i < 3 {
		fmt.Println("goto", i)
		i++
		goto inicio
	}
	fmt.Println("fim")
}
