package main

import "fmt"

func main() {
	s := "banana"

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				fmt.Println("sem rótulo:", i, j)
				break
			}
		}
	}

externo:
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				fmt.Println("com rótulo", i, j)
				continue externo
			}
		}
	}
}
