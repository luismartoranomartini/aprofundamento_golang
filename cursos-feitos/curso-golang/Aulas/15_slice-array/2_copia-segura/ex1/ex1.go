package main

import "fmt"

// cópia simples
func main() {
	fmt.Println("---cópia insegura---")
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = slice1

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = 10
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice2 = append(slice2, 6)
	slice2[0] = 99 // perdi a referência do arry interno
	fmt.Println(slice1)
	fmt.Println(slice2)

	// função nativa - copy()
	fmt.Println("---cópia segura---")
	var slice3 = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, len(slice3))
	_ = copy(slice4, slice3)
	fmt.Println("Slice 3:", slice3)
	fmt.Println("Slice 4:", slice4)
	fmt.Println("---alterado---")
	slice3[0] = 10

	fmt.Println("Slice 3:", slice3)
	fmt.Println("Slice 4:", slice4)
}
