package main

import "fmt"

func main() {
	for i := 1; i <= 6; i++ {
		if i == 2 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Print(i, ", ")
	}
}
