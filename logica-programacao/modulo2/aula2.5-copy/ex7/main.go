package main

import (
	"fmt"
	"runtime"
)

func heapAlloc() uint64 {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc
}

func main() {
	grande := make([]int, 1_000_000)
	copia := make([]int, 3)
	copy(copia, grande[:3])
	grande = nil

	fmt.Println("B: heap acima de 7 MB?", heapAlloc() > 7_000_000)
}
