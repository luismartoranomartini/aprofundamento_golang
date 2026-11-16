package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	bad := "a\xffb"

	fmt.Println(len(bad))                 // 3  (a, \xff, b: 3 bytes)
	fmt.Println(utf8.ValidString(bad))    // false
	fmt.Println(utf8.ValidString("ação")) // true

	for i, r := range bad {
		fmt.Printf("%d:%U \n", i, r)
	}

	r := []rune(bad)
	volta := string(r)
	fmt.Println(len(volta), volta == bad)
	fmt.Printf("%q\n", volta)

	// fatiar ã ao meio
	meio := "ã"[:1]
	fmt.Printf("%q\n", meio)
	fmt.Println(utf8.ValidString(meio))

	s := "\xc3"
	n := 0
	for range s {
		n++
	}
	fmt.Println(n, len(s), utf8.ValidString(s))
}
