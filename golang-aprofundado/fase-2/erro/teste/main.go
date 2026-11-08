package main

import (
	"errors"
	"fmt"
)

var ErrA = errors.New("falhou")

func f(n int) error {
	if n == 0 {
		return ErrA
	}
	if n == 1 {
		return errors.New("falhou")
	}
	return nil
}

func main() {
	fmt.Println(f(0) == ErrA, f(1) == ErrA, f(2) == ErrA)
	fmt.Println(f(1), f(2))
	fmt.Println(f(1) == f(1))
}
