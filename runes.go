package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println("s", s)
	fmt.Println("Len s", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Run count:", utf8.RuneCountInString(s))

	for i, runeValue := range s {
		fmt.Printf("%#U start at %d", runeValue, i)
	}

	fmt.Println("\nUsing DecodeRuneInString")

	// Decode Rune In String
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U start at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found t")
	} else if r == 'ส' {
		fmt.Println("found ส")
	}
}
