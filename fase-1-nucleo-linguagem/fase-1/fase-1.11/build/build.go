package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder

	for i := 1; i <= 3; i++ {
		fmt.Fprintf(&sb, "item%d;", i)
	}
	sb.WriteString("fim")
	sb.WriteByte('!')
	sb.WriteRune('ç')
	fmt.Println(sb.String(), sb.Len())

	s := "casa"
	r := []rune(s)
	r[0] = 'm'
	fmt.Println(s)
	fmt.Println(string(r))
}
