package main

import "fmt"

func main() {
	b := []byte("ir para casa")
	for i := range b {
		if b[i] == ' ' {
			b[i] = '_'
		}
	}
	fmt.Println(b)
	fmt.Println(string(b))
}
