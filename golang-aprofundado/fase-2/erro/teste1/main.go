package main

import (
	"errors"
	"fmt"
)

var ErrA = errors.New("A")

type EB struct{ N int }

func (e *EB) Error() string { return fmt.Sprint("EB", e.N) }

func main() {
	e1 := fmt.Errorf("c1: %w", ErrA)
	e2 := fmt.Errorf("c2: %v", ErrA)
	e3 := fmt.Errorf("c3: %w", &EB{7})
	e4 := fmt.Errorf("c4: %w", e3)

	fmt.Println(errors.Is(e1, ErrA), errors.Is(e2, ErrA))
	var eb *EB
	fmt.Println(errors.As(e4, &eb), eb.N)
	fmt.Println(e4)
	fmt.Println(e1 == ErrA, errors.Unwrap(e1) == ErrA)
}
