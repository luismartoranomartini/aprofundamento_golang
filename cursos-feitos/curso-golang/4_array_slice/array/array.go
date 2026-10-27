package main

import "fmt"

func main() {
	fmt.Println("ARRAYS & SLICES")

	var array1 [5]int
	array1[0] = 55

	fmt.Println(array1)

	array2 := [5]string{"p1", "p2", "p3", "p4", "p5"}
	// array2[6] = "p6"
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	// array3[6] = 6
	fmt.Println(array3)
}
