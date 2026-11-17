package main

import "fmt"

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	t1 := s1[:2]
	t1 = append(t1, 99)
	fmt.Println("sem limites: s =", s1, "| t =", t1)
	fmt.Println("mesmo array (s1, s2)?", &s1[0] == &t1[0])

	s2 := []int{10, 20, 30, 40, 50}
	t2 := s2[:2:2]
	t2 = append(t2, 99)
	fmt.Println("com limite: s=", s2, "| t =", t2)
	fmt.Println("mesmo array (s2, t2)", &s2[0] == &t2[0])
}
