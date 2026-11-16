package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(utf8.RuneLen('a'), utf8.RuneLen('ç'), utf8.RuneLen('😀'))
}
