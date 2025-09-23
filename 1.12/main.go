package main

import(
	"log"
)

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for _, str := range strs {
		set[str] = struct{}{}
	}

	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}	

	log.Println(res)
}
