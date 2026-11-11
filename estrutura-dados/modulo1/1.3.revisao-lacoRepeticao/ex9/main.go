package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
