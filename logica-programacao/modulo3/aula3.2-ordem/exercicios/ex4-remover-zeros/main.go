package main

import "fmt"

func removerZeros(m map[string]int) {

	for k, v := range m {
		if v == 0 {
			delete(m, k)
		}
	}

}

func main() {
	mapa := map[string]int{"a": 0, "b": 1, "c": 0, "d": 2}

	fmt.Println(len(mapa))
	removerZeros(mapa)
	fmt.Println(mapa)
	fmt.Println(len(mapa))
}
