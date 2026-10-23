package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var v struct{}
	fmt.Println("tamanho de struct{}:", unsafe.Sizeof(v))
	var i int64
	fmt.Println("tamanho de int64", unsafe.Sizeof(i))

	visto := map[string]struct{}{}
	visto["a"] = struct{}{}
	visto["b"] = struct{}{}
	visto["a"] = struct{}{}
	fmt.Println(len(visto))
	_, ok := visto["a"]
	_, ok2 := visto["z"]
	fmt.Println(ok, ok2)
}
