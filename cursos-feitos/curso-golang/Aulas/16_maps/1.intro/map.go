package main

import (
	"fmt"
	"maps"
	"slices"
)

// maps lidam com concorrência
// chave - valor
func main() {
	// var meuMap = map[string]any{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// 	"d": 4,
	// 	"e": 5,
	// 	"f": 6,
	// }
	//
	// var chaves []string
	//
	// for chave := range meuMap {
	// 	chaves = append(chaves, chave)
	// }
	// sort.Strings(chaves)
	// fmt.Println(chaves)
	//
	// // percorrer a lista já ordenada
	// for _, chave := range chaves {
	// 	valor := meuMap[chave]
	// 	fmt.Println(chave, valor)
	//
	var meuMap = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}
	for _, chave := range slices.Sorted(maps.Keys(meuMap)) {
		fmt.Println(chave, meuMap[chave])
	}
}
