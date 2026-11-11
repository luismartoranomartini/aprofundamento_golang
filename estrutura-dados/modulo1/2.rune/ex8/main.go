package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	for _, r := range "aé世😀" {
		fmt.Printf("%c U+%04X %d bytes\n", r, r, utf8.RuneLen(r))
	}
}
