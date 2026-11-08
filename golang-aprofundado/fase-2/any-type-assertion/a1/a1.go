package main

import "fmt"

func main() {
	var x any = 42
	fmt.Printf("%T %v\n", x, x)
	x = "texto"
	fmt.Printf("%T %v\n", x, x)
	x = []int{1, 2}
	fmt.Printf("%T %v\n", x, x)

	var y any = 10
	n, ok := y.(int)
	fmt.Println(n+1, ok)
}
