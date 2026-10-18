package main

import "fmt"

func Ruim(yield func(int) bool) {
	for i := 1; i <= 3; i++ {
		yield(i)
	}
}

func main() {
	for v := range Ruim {
		fmt.Println(v)
		break
	}
}
