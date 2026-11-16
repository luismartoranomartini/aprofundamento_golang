package main

import "fmt"

func main() {
	b := make([]int, 0, 5)
	fmt.Println(b[0:1], len(b[0:1]))
	fmt.Println(b[1:])
}
