package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "banana"
	contA := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			contA++
		}
	}
	fmt.Println(contA)

	contB := strings.Count(s, "a")
	fmt.Println("laço:", contB)
}
