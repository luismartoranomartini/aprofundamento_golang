package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		ErrA = errors.New("erro A")
		ErrB = errors.New("erro B")
	)

	j := errors.Join(ErrA, nil, ErrB)
	fmt.Println(j)
	fmt.Println(errors.Is(j, ErrA), errors.Is(j, ErrB))
	fmt.Println(errors.Join(nil, nil) == nil)

	multi := fmt.Errorf("duas causas: %w e %w", ErrA, ErrB)
	fmt.Println(multi)
	fmt.Println(errors.Is(multi, ErrA), errors.Is(multi, ErrB))
}
