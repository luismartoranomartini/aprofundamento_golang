package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}
	y := x
	y[2] = 9
	fmt.Println(x, y)

}
