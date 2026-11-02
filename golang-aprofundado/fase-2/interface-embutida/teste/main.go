package main

import "fmt"

type T struct{ N int }

func (t *T) String() string { return fmt.Sprintf("T(%d)", t.N) }

func main() {
	t := T{1}
	fmt.Println(t)
	fmt.Println(&t)
	var s fmt.Stringer = &t
	t.N = 5
	fmt.Println(s)
}
