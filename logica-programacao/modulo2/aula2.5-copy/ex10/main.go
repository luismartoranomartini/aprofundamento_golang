package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	max := 6
	fmt.Println(s[1:2:max])
}
