package main

import "fmt"

func main() {
	x := 1
	f := func() { fmt.Println("x =", x) }
	x = 2
	f()

	g := func() { x = 100 }
	g()
	fmt.Println(x)

	total := 0
	soma := func(n int) { total += n }
	ver := func() int { return total }
	soma(5)
	soma(7)
	fmt.Println(ver())
}
