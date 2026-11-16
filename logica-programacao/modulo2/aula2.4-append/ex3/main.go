package main

import "fmt"

func main() {
	var s []int
	antes := cap(s)

	for i := 0; i < 2000; i++ {
		s = append(s, i)
		if cap(s) != antes {
			if antes == 0 {
				fmt.Printf("len=%4d cap: %4d -> %4d\n", len(s), antes, cap(s))
			} else {
				fmt.Printf("len=%4d cap: %4d -> %4d (fator %.2f)\n", len(s), antes, cap(s), float64(cap(s))/float64(antes))
			}
			antes = cap(s)
		}
	}
}
