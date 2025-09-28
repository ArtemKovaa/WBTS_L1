package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Go(func() { work() })
	wg.Go(func() { work() })
	wg.Go(func() { work() })

	wg.Wait()
}

func work() {
	log.Println("Before sleep")
	sleep(10)
	log.Println("After sleep")
}


func sleep(sec int) {
	select {
	case <-time.After(time.Duration(sec) * time.Second):
		return
	}
}
