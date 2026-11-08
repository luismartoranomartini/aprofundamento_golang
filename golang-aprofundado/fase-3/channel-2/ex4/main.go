package main

import "fmt"

func main() {
	a := make(chan int, 1)
	b := make(chan int, 1)
	contA, contB := 0, 0

	for i := 0; i < 1000; i++ {
		a <- 1
		b <- 1
		select {
		case <-a:
			contA++
			<-b
		case <-b:
			contB++
			<-a
		}
	}
	fmt.Println("a:", contA, "b:", contB)
}
