package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s[2:5])
	fmt.Println(s[:3])
	fmt.Println(s[4:])
	t := s[1:4]
	u := t[1:3]
	fmt.Println(t, u)
	fmt.Println(u[0])
	u[0] = 0
	fmt.Println(s)

}
