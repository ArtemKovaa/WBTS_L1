package main

import "fmt"

func main() {
	s := "главрыба"

	fmt.Println(reverse(s))
}

func reverse(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
