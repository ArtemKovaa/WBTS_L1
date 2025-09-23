package main

import (
	"log"
)

func main() {
	a := []int{1, 11, 111, 0, 0, 0, 2, 3}
	b := []int{2, 3, 4, 0, 0, -10, 111, 23013012, 11, 32, 0}

	set := make(map[int]struct{})
	var exists struct{}

	for _, v := range a {
		set[v] = exists
	}

	res := []int{}

	for _, v := range b {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}

	log.Print(res)
}