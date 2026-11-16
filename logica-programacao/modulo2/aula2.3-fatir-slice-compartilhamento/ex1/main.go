package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	fmt.Println(s[1:4], len(s[1:4]), cap(s[1:4]))
	fmt.Println(s[:2], len(s[:2]), cap(s[:2]))
	fmt.Println(s[2:], len(s[2:]), cap(s[2:]))
	fmt.Println(s[:], len(s[:]), len(s[:]))
	fmt.Println(s[2:2], len(s[2:2]), cap(s[2:2]))
}
