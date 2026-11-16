package main

import "fmt"

func main() {
	s := "Go"
	i := 5
	fmt.Println(s[i])
	// panic: runtime error: index out of range [5] with length 2
}
