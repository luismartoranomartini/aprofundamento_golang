package main

import "fmt"

func alterarMap(mapa map[string]int) { // passa cópia do objeto na memória
	// map sempre passa o endereço
	fmt.Printf("Dentro mapa: %p\n", &mapa)
	mapa["Luís"] = 10
}

func alteraInt(a int) {
	a = 20
	fmt.Printf("Dentro: %p\n", &a)
}

func main() {
	a := 10
	mapa := map[string]int{"Luís": 1}
	fmt.Printf("Fora mapa: %p\n", &mapa)
	alterarMap(mapa)
	fmt.Printf("fora: %p\n", &a)
	alteraInt(a)
	println(mapa["Luís"])
}
