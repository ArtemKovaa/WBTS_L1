package main

import (
	"math/rand"
	"sync"
	"log"
	"time"
)

var m = make(map[int]int)
var mutex = sync.RWMutex{}


func main() {	
	go worker_1()
	go worker_2()

	time.Sleep(10 * time.Second)
}

func worker_1() {
	for {
		k := rand.Intn(100_000)
		v := rand.Intn(100_000)
		mutex.RLock()
		m[k] = v
		log.Printf("Set value %d for key %d", v, k)
		mutex.RUnlock()
		time.Sleep(100 * time.Millisecond)
	}
}

func worker_2() {
	for {
		k := rand.Intn(100_000)
		mutex.Lock()
		v, ok := m[k]
		if !ok {
			log.Println("Value not found")
		} else {
			log.Printf("Fonud value %d by key %d", v, k)
		}
		mutex.Unlock()
	}
}
