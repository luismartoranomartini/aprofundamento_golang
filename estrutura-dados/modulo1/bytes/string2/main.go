package main

import "fmt"

func main() {
	s := "Gopher"
	for i := 0; i < len(s); i++ {
		fmt.Printf("i:=%d byte=%d char=%c\n", i, s[i], s[i])
	}
}
