package main

import (
	"fmt"
)

func removeVowels(word string) (string, string) {
	var vowels string
	var other string
	for _, rune := range word {
		switch rune {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels += string(rune)
		default:
			other += string(rune)
		}
	}
	return other, vowels
}

func main() {
	fmt.Println(removeVowels("aeiou"))               // returns "", "aeiou"
	fmt.Println(removeVowels("Apple"))               // returns "ppl", "Ae"
	fmt.Println(removeVowels("Waffles are awesome")) // returns "Wffls r wsm", "aeaeaeoe"
}
