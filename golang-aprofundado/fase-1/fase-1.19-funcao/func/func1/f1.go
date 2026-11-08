package main

import (
	"errors"
	"fmt"
)

func soma(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

func minMax(xs []int) (min, max int) {
	min, max = xs[0], xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return
}

func main() {

	fmt.Println(soma(2, 3))

	q, err := divide(10, 2)
	fmt.Println(q, err)
	_, err = divide(1, 0)
	fmt.Println(err)

	lo, hi := minMax([]int{4, 9, 1, 7})
	fmt.Println(lo, hi)

}
