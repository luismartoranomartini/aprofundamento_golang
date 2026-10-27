package main

import "fmt"

func main() {

	array := [2]int{2, 4}
	fmt.Printf("%T\n", array)

	soma := 0
	for i := 0; i < len(array); i++ {
		soma += array[i]
	}

	fmt.Println(soma)

	somaTotal := array[0] + array[1]
	fmt.Println(somaTotal)

}
