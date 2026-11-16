package main

import "fmt"

func main() {

	soma := 0

	for i := 0; i <= 100; i++ {
		soma += i
	}
	fmt.Println(soma)
}
