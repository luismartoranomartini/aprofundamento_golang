package main

import "fmt"

func main() {
	s := "olá"
	fmt.Printf("%q\n", s[:3])
	fmt.Printf("%q\n", string([]rune(s)[:3]))
}
