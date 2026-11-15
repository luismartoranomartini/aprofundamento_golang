package main

import "fmt"

func main() {
	s := "banana"
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				fmt.Println(i, j)
				goto fim
			}
		}
	}
fim:
	fmt.Println("terminou")
}
