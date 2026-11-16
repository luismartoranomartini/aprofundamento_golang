package main

import "fmt"

func main() {
	s := " Martini "
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(s[:i])
	}
	// for i := 0; i < len(s); i-- {
	// 	fmt.Println(s[:i])
	// }
}
