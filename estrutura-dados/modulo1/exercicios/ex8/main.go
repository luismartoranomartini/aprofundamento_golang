package main

import "fmt"

func mesmaPonta(s string) bool {
	if s == "" {
		return false
	}
	return s[0] == s[len(s)-1]

}

func main() {
	sim := []string{"abca", "xx", "g", "otto", "b12b"}
	nao := []string{"ab", "gato", "abcd", "12"}

	for _, palavra := range sim {
		fmt.Printf("%-6q -> %t\n", palavra, mesmaPonta(palavra))
	}
	for _, palavra := range nao {
		fmt.Printf("%-6q -> %t\n", palavra, mesmaPonta(palavra))
	}
}
