package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "привет"
	fmt.Println("s:", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}

	fmt.Println("len(s):", len(s))
	fmt.Println("runecount(s):", utf8.RuneCountInString(s))

	for idx, rune := range s {
		fmt.Printf("%#U starts at %d\n", rune, idx)
	}

}
