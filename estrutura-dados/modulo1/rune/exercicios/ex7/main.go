package main

import "fmt"

func ehPalindromo(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		j := len(r) - 1 - i
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func main() {
	s := "1aç"
	fmt.Println(ehPalindromo(s))
}
