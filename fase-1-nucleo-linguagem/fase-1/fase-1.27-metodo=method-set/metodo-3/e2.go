package main

type Contador struct {
	N int
}

func (c *Contador) Inc() { c.N++ }

func main() {
	m := map[string]Contador{"a": {}}
	m["a"].Inc()
}
