func cresce(s []int) {
	s = append(s, 7)
}

func main() {
	s := make([]int, 2, 4)
	cresce(s)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s[:3]) // olhando além do len
}
