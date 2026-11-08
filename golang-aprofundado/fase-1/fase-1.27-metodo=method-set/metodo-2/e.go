package main

type Contador struct {
	N int
}

func (c Contador) IncV()    { c.N++ }
func (c *Contador) inc()    { c.N++ }
func (c Contador) Ver() int { return c.N }

func main() {
	Contador{}.Inc()
}
