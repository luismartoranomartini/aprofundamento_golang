package main

import "fmt"

func main() {
	s := "Go123"

	for i := 0; i < len(s); i++ {
		b := s[i]
		fmt.Printf("pos=%d byte=%d char=%c tipo=%T\n", i, b, b, b)
	}
}
