package main

import "fmt"

func produzir(valores ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range valores {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func main() {
	a := produzir(1, 2, 3)
	b := produzir(10, 20)

	soma := 0
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				continue
			}
			soma += v
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			soma += v
		}
	}
	fmt.Println("soma:", soma)
}
