package main

import "fmt"

func main() {
	slice := make([]float64, 10, 11)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 5)
	slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := make([]float32, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2 = append(slice2, 6)
	slice2 = append(slice2, 7)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

}
