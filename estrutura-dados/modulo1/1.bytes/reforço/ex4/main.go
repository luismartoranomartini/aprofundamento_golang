package main

import "fmt"

func main() {
	s := "abc123def45"
	contDigito := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			contDigito++
		}
	}
	fmt.Println(contDigito)
}
