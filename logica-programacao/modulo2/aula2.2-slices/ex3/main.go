package main

import "fmt"

func main() {
	s := []int{10, 20, 30}
	for i := 0; i < len(s); i++ {
		fmt.Print(i, ":", s[i], " ")
	}
	fmt.Println()
	for i, v := range s {
		fmt.Print(i, "=", v, " ")
	}
	fmt.Println()
	s[1] = 99
	fmt.Println(s)
}
