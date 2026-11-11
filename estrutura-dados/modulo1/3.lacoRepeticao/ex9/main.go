package main

import "fmt"

func main() {
	b := []byte("go")
	for _, c := range b {
		c -= 32
		fmt.Printf("%c ", c)
	}
	fmt.Println(string(b))

	for i := range b {
		b[i] -= 32
	}
	fmt.Println(string(b))
}
