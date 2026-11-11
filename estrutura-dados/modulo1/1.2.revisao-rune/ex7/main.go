package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%c %c", unicode.ToUpper('ç'), unicode.ToLower('Â'))
}
