package main

import "fmt"

func main() {
	original := []int{1, 2, 3}

	c1 := make([]int, len(original))
	copy(c1, original)

	c2 := append([]int(nil), original...)

	c1[0] = 100
	c2[1] = 200

	fmt.Println("original:", original)
	fmt.Println("c1: ", c1)
	fmt.Println("c2: ", c2)
	fmt.Println("mesmo array (original, c1)?", &original[0] == &c1[0])
	fmt.Println("mesmo array (original, c2)?", &original[0] == &c2[0])

	vazio := []int{}
	fmt.Println("clone de vazio por append é nil?", append([]int(nil), vazio...) == nil)
	cv := make([]int, len(vazio))
	copy(cv, vazio)
	fmt.Println("clone de vazio por make é nil?", cv == nil)
}
