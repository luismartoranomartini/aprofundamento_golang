package main

import "fmt"

type Numero interface {
	~int | ~float64
}

type Celsius float64

func Soma[T Numero](s []T) T {
	var t T
	for _, v := range s {
		t += v
	}
	return t
}

func main() {
	fmt.Println(Soma([]int{1, 2, 3}))
	fmt.Println(Soma([]float64{1.5, 2.5}))
	t := Soma([]Celsius{10.5, 20})
	fmt.Printf("%v %T\n", t, t)
}
