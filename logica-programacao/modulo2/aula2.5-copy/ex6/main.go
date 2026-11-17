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
	pequeno := grande[:3]
	grande = nil
	fmt.Println("A: heap acima de 7 MB", heapAlloc() > 7_000_000, pequeno)
}
