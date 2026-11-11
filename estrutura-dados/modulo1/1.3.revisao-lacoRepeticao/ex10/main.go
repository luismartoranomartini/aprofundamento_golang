package main

import "fmt"

func main() {
	s := "banana"
	for i := 0; i < len(s); i++ {
		if s[i] == 'n' {
			fmt.Printf("A letra %c está na %dª posição", s[i], i)
			break
		}
	}
}
