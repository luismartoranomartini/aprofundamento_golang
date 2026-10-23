package main

import "fmt"

func somador(base int) func(int) int {
	return func(n int) int { return base + n }
}

func main() {
	soma10 := somador(10)
	soma100 := somador(100)
	fmt.Println(soma10(1), soma100(1))

	var fs []func() int
	for i := 0; i < 3; i++ {
		fs = append(fs, func() int {
			i += 10
			return i
		})
	}
	for _, f := range fs {
		fmt.Print(f(), " ", f(), " | ")
	}
	fmt.Println()
}
