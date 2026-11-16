package main

import "fmt"

func main() {
	p := 1
	for range 10 {
		p *= 2
	}
	fmt.Println(p)
}
