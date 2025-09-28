package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isUniqueRunes("abcd"))
	fmt.Println(isUniqueRunes("abCdefAaf"))
	fmt.Println(isUniqueRunes("aabcd"))
	fmt.Println(isUniqueRunes("5324абвдфы"))
	fmt.Println(isUniqueRunes("Тест"))
}

func isUniqueRunes(s string) bool {
	set := make(map[rune]struct{})
	for _, r := range s {
		lowerR := unicode.ToLower(r)
		if _, ok := set[lowerR]; ok {
			return false
		}
		set[lowerR] = struct{}{}
	}
	return true
}
