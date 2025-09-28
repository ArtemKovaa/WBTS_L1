package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"

	fmt.Println(reverse(s))
}

func reverse(s string) string {
	words := strings.Split(s, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
