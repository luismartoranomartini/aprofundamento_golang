package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "abc"
	r := 'á'
	rs := []rune(s)
	s2 := string(rs)
	n := len(rs)
	n = utf8.RuneCountInString(s)
	for i, r := range s {
		fmt.Println(i, r)
	}
	utf8.RuneLen(r)

}
