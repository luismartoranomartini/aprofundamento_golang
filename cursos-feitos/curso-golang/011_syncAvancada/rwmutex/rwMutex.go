package main

import (
	"fmt"
	"sync"
	"time"
)

// RwMutex é mais flecível

type Cache struct {
	dados map[string]any
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		dados: make(map[string]any),
	}
}

func (c *Cache) Set(chave string, valor any) {
	logf("Set: tentativa de Escrita (Lock)")
	c.mu.Lock()
	logf("Set: Lock Adquirida")
	time.Sleep(time.Millisecond * 300)

	c.dados[chave] = valor
	c.mu.Unlock()

	logf("Set: Unlock (escrita finalizada)")

}

func (c *Cache) Get(chave string) (any, bool) {
	logf("Get: tentativa de Leitura (Rlock)")
	c.mu.RLock()
	logf("Get: tentativa de Leitura Adquirida")
	time.Sleep(time.Millisecond * 100)
	c.mu.RUnlock()
	valor, ok := c.dados[chave]
	logf("Get: RUnlock (Leitura finalizada)")
	return valor, ok
}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[%4dms] - Mensagem [%s]\n", time.Since(start).Milliseconds(), msg)
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)
	wg.Add(4)
	cache.dados["k"] = "v"
	go func() {
		defer wg.Done()
		logf("Goroutine 1: chamando Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 1: recebeu retorno do Get()")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 10)
		logf("Goroutine 2: chamando Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 2: recebeu retorno do Get()")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 20)
		logf("Goroutine 3: chama Set() - (Quer lock)")
		cache.Set("k", "v2")
		logf("Goroutine 3: recebeu retorno do Set()")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 300)
		logf("Goroutine 4: chamando Get()")
		_, _ = cache.Get("k")
		logf("Goroutine 4: recebeu retorno do Get()")
	}()
	wg.Wait()

	logf("Fim")
}
