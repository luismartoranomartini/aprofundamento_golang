package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(10)
	generica(false)

	mapa := map[interface{}]interface{}{
		1:      "string",
		"nome": 2,
		true:   float64(12),
	}

	fmt.Println(mapa)
}
