package main

import "fmt"

func main() {
	nfc := "\u00e9"  // é como 1 ponto de código: U+00E9
	nfd := "e\u0301" // é como 2: 'e' + acento combinante U+0301

	fmt.Println(nfc, nfd) // é é   (parecem iguais na tela)

	fmt.Println(len(nfc), len([]rune(nfc))) // 2 1   (2 bytes, 1 rune)
	fmt.Println(len(nfd), len([]rune(nfd))) // 3 2   (3 bytes, 2 runes)
	fmt.Println(nfc == nfd)                 // false (os bytes diferem)

	for i, r := range nfd {
		fmt.Printf("%d:%U \n", i, r)
	}
	// 0:U+0065 1:U+0301

	//emoji
	br := "🇧🇷"
	fmt.Println(len(br), len([]rune(br))) // 8 2   (uma bandeira = 2 runes)

	fam := "👨‍👩‍👧"
	fmt.Println(len(fam), len([]rune(fam))) // 18 5  (um emoji = 5 runes)
	// os 5: U+1F468 U+200D U+1F469 U+200D U+1F467  (os U+200D juntam as figuras)

	a := "e\u0301"
	fmt.Println(len(a), len([]rune(a)))
}
