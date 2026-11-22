package main

import (
	"fmt"
)

func main() {
	s := "São Paulo, coração"

	count := 0
	for _, r := range s {
		if r > 127 {
			count++
		}
	}
	fmt.Println(count)

}
