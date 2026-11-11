package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "olá"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))
	fmt.Println(utf8.RuneCountInString(s))
}
