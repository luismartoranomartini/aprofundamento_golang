package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrNegativo = errors.New("número negativo")

func raiz(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativo
	}
	r := 0
	for (r+1)*(r+1) <= n {
		r++
	}
	return r, nil
}

func main() {
	v, err := strconv.Atoi("42")
	fmt.Println(v, err)
	v, err = strconv.Atoi("abc")
	fmt.Println(v, err)

	for _, n := range []int{16, -4} {
		r, err := raiz(n)
		if err == ErrNegativo {
			fmt.Println("n =", n, "-> erro sentinha:", err)
			continue
		}
		if err != nil {
			fmt.Println("outro erro:", err)
			continue
		}
		fmt.Println("n =", n, "-> raiz", r)
	}
}
