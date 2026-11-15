package main

import "fmt"

func main() {
	i := 0
inicio:
	if i < 3 {
		fmt.Println(i, " ")
		i++
		goto inicio
	}
}
