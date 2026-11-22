package main

import "fmt"

func ehPalindromo(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	s := "arara"
	nome := "gato"
	n := "abca"
	v := ""
	x := "x"
	d := "çaç"

	fmt.Println(ehPalindromo(s))
	fmt.Println(ehPalindromo(nome))
	fmt.Println(ehPalindromo(n))
	fmt.Println(ehPalindromo(v))
	fmt.Println(ehPalindromo(x))
	fmt.Println(ehPalindromo(d))
}
