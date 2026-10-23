package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		switch {
		case i == 1:
			fmt.Println("um")
			fallthrough
		case i == 2:
			fmt.Println("dois")
		case i == 3:
			continue
		default:
			fmt.Println("outro", i)
		}
		fmt.Println("fim da volta", i)
	}
}
