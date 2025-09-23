package main

import "log"

func main() {
	a := 10 // 01
	b := 99 // 10

	// a, b = b, a - also working
	a, b = (a ^ b) ^ a, (a ^ b) ^ b

	log.Printf("a = %d", a)
	log.Printf("b = %d", b)
}
