package main

import (
	"fmt"
	"sync"
)

func main() {
	mtx := sync.Mutex{}
	wg := sync.WaitGroup{}
	arr := [5]int{2, 4, 6, 8, 10}

	for _, v := range arr {
		wg.Go(
			func() { 
				mtx.Lock()
				fmt.Println(v * v)
				mtx.Unlock()
			},
		)
	}
	wg.Wait()
}