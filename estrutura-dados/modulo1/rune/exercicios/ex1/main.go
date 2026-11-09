package main

import "fmt"

func main() {
	s := "coração"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
	fmt.Println(len(s) == len([]rune(s)))
}
