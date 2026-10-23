package main

import (
	"fmt"
	"runtime"
)

func tenta(nome string, f func()) {
	defer func() {
		r := recover()
		if e, ok := r.(runtime.Error); ok {
			fmt.Println(nome, "-> runtime.Error:", e)
		} else {
			fmt.Println(nome, "->", r)
		}
	}()
	f()
}

func main() {
	tenta("índice", func() { s := []int{1}; i := 5; _ = s[i] })
	tenta("divisão", func() { a, b := 1, 0; _ = a / b })
	tenta("mapa nil", func() { var m map[string]int; m["x"] = 1 })
	tenta("nosso", func() { panic("meu erro") })
	tenta("sem panic", func() {})
	fmt.Println(recover())
}
