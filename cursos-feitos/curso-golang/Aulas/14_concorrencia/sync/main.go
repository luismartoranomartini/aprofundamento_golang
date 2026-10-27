package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	callDataBase(&wg)
	callAPI(&wg)
	processInternal(&wg)

	wg.Wait()

}
func callDataBase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizado callDataBase")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado callAPI")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizado processInternal")
	wg.Done()
}