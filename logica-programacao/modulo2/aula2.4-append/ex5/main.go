package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	t := s[:2]
	fmt.Printf("t=%v len=%d cap=%d\n", t, len(t), cap(t))

	t = append(t, 99)

	fmt.Println("s =", s)
	fmt.Println("t =", t)
}
