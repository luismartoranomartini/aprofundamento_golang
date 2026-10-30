package main

import (
	"fmt"
)

func main() {
	numTotal := 5372

	milhar := numTotal / 1000
	restoMilhar := numTotal % 1000
	centena := restoMilhar / 100
	restoCentena := restoMilhar % 100
	dezena := restoCentena / 10
	restoDezena := restoCentena % 10
	unidade := restoDezena

	fmt.Println(milhar)
	fmt.Println(centena)
	fmt.Println(dezena)
	fmt.Println(unidade)

	fmt.Println(milhar + centena + dezena + unidade)
}
