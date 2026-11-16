package main

import "fmt"

func main() {
	soma := 0
	for i := 1; i <= 9; i += 2 {
		soma += i
	}
	fmt.Println(soma)
}
