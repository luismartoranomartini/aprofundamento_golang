package main

import "fmt"

func main() {
	s := "Go"
	fmt.Println(len(s))
	fmt.Println(s[0], s[1])
	fmt.Printf("%c %c\n", s[0], s[1])
	fmt.Printf("%T\n", s[0])
}
