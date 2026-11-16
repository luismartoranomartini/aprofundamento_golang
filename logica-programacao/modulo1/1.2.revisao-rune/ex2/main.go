package main

import "fmt"

func main() {
	for i, r := range "ação" {
		fmt.Printf("%d: %c | ", i, r)
	}
}
