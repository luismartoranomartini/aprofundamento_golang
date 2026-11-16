package main

import "fmt"

func main() {
	n := 1
	for {
		if n > 100 {
			break
		}
		n *= 2
	}
	fmt.Println(n)
}
