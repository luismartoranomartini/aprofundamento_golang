package main

import "fmt"

func main() {
	s := "abcdef"
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("%c-%c ", s[i], s[j])
	}
}
