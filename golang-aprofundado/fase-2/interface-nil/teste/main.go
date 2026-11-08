package main

import "fmt"

type E struct{}

func (*E) Error() string { return "E" }

func f(n int) error {
	var e *E
	if n > 0 {
		e = &E{}
	}
	return e
}

func g(n int) error {
	if n > 0 {
		return &E{}
	}
	return nil
}

func main() {
	fmt.Println(f(0) == nil, f(1) == nil)
	fmt.Println(g(0) == nil, g(1) == nil)
	var p *E
	var err error = p
	fmt.Println(p == nil, err == nil)
}
