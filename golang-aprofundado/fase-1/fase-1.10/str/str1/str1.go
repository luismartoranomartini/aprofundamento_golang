package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "ação"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len([]rune(s)))

	fmt.Println(s[0], s[1], s[2])
	fmt.Println("%x\n", s)

	for i, r := range s {
		fmt.Println(i, r, string(r))
	}
}
