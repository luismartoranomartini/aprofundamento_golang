package main

import "fmt"

func main() {
	var r rune = 'a'
	var b byte = 'a'
	fmt.Println(r == rune(b))
	fmt.Println(b == byte(r))
}
