package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 10)
	fmt.Println("Slice1", slice)
	fmt.Println(reflect.TypeOf(slice))

	array := [5]int{6, 7, 8, 9, 10}
	fmt.Println("Array", array)

	// slice2 := array[1:3]
	// fmt.Println(slice2)
	// fmt.Println(reflect.TypeOf(array))

	array[1] = 5
	fmt.Println("Slice mudança(?)", slice)
	fmt.Println("Array mudança(?)", array)
}
