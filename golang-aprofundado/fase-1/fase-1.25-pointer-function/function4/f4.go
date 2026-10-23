package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	p := &s[0]
	*p = 100
	fmt.Println(s)

	fmt.Println("len/cap antes", len(s), cap(s))
	s = append(s, 4)
	fmt.Println("len/cap depois:", len(s), cap(s))
	*p = 999
	fmt.Println("s:", s, "| *p:", *p)
}
