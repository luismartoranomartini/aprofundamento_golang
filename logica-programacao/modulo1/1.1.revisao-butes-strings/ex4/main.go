package main

import "fmt"

func main() {
	s := "Go Lang Rust"
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			n++
		}
	}
	fmt.Println(n)
}
