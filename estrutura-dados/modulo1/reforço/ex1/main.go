package main

import "fmt"

func main() {
	s := "Compilador"

	primeiro := s[0]
	ultimo := s[len(s)-1]

	fmt.Printf("len=%d primeiro=%c (%d) último=%c (%d)\n",
		len(s),
		primeiro, primeiro,
		ultimo, ultimo,
	)
}
