package main

import "fmt"

type P struct {
	N int
	T []int
}

func f(p P)  { p.N = 10; p.T[0] = 10; p.T = append(p.T, 7) }
func g(p *P) { p.N = 20; p.T[1] = 20; p.T = append(p.T, 8) }

func main() {
	x := P{1, []int{1, 2}}
	f(x)
	fmt.Println(x.N, x.T)
	g(&x)
	fmt.Println(x.N, x.T)
}
