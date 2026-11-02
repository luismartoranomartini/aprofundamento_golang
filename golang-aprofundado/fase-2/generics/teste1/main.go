package main

import (
	"fmt"
	"strconv"
)

type List[E any] []E

func (l List[E]) Apply[F any](f func(E) F) List[F] {
	r := make(List[F], len(l))
	for i, x := range l {
		r[i] = f(x)
	}
	return r
}

func main() {
	l := List[int]{1, 2, 3}
	s := l.Apply(func(n int) string { return strconv.Itoa(n) + "!" })
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}
