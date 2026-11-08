package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	// 1) clousure guardadas e chamadas depois
	var fs []func() int
	for i := 0; i < 3; i++ {
		fs = append(fs, func() int { return i })
	}
	for _, f := range fs {
		fmt.Print(f(), " ")
	}
	fmt.Println()

	// 2) endereços de i
	var ps []*int
	for i := 0; i < 3; i++ {
		ps = append(ps, &i)
	}
	fmt.Println(ps[0] == ps[1], ps[1] == ps[2])
	fmt.Println(*ps[0], *ps[1], *ps[2])

	// 3) goroutines
	var wg sync.WaitGroup
	var mu sync.Mutex
	var res []int
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			res = append(res, i)
			mu.Unlock()
		}()
	}
	wg.Wait()
	sort.Ints(res)
	fmt.Println(res)

	// 4) range também
	var gs []func() string
	for _, s := range []string{"a", "b", "c"} {
		gs = append(gs, func() string { return s })
	}
	for _, g := range gs {
		fmt.Print(g(), " ")
	}
	fmt.Println()

}
