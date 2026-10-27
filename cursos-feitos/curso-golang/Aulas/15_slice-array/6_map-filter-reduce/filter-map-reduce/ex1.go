package main

import "fmt"

type mySlice []int

func main() {
	var lista = mySlice{1, 2, 3, 4, 5, 6}
	lista = lista.Filter(func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(lista)

	lista = lista.Map(func(i int) int {
		return i * 2
	})
	fmt.Println(lista)

	soma := lista.Reduce(func(i1, i2 int) int {
		return i1 * i2
	}, 1)
	fmt.Println(soma)
	fmt.Println("-------encadeamento------")

	var lista2 = mySlice{1, 2, 3, 4, 5, 6}
	mult := lista2.Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i * 2
	}).Reduce(func(i1, i2 int) int {
		return i1 * i2
	}, 1)
	fmt.Println(mult)
}

// encadeamento de métodos
func (m mySlice) Filter(cond func(int) bool) mySlice {
	var result mySlice

	for _, numero := range m {
		if cond(numero) {
			result = append(result, numero)
		}
	}
	return result
}

func (m mySlice) Map(trans func(int) int) mySlice {
	var result mySlice

	for _, numero := range m {
		result = append(result, trans(numero))
	}
	return result
}

func (m mySlice) Reduce(acc func(int, int) int, inicial int) int {
	var result = inicial
	for _, numero := range m {
		result = acc(result, numero)
	}
	return result
}
