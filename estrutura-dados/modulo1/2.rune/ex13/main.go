package main

import (
	"fmt"
	"unicode"
)

func main() {
	r := []rune("éric")
	r[0] = unicode.ToUpper(r[0])
	fmt.Println(string(r))
	fmt.Println(unicode.IsLetter('ç'))
	fmt.Println(unicode.IsDigit('7'))
}
