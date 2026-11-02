package main

import "fmt"

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type MyInt int

func add[T Integer](a, b T) T { return a + b }

func main() {

	fmt.Println(add(MyInt(1), MyInt(2)))
}
