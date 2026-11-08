package main

import "fmt"

type C struct{ N int }

func (c C) A()  { c.N += 1 }
func (c *C) B() { c.N += 10 }

func main() {
	x := C{}
	x.A()
	x.B()
	p := &x
	p.A()
	p.B()
	fmt.Println(x.N)
	y := x
	y.B()
	fmt.Println(x.N, y.N)
}
