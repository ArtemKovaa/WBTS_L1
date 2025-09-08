package main

import (
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
)

func work(id int, ch <-chan int) {
	for msg := range ch {
		log.Println("Worker", id, "read message:", msg)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Enter number of workers as first argument")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Incorrect number of workers")
	}

	ch := make(chan int)

	for i := range n {
		go work(i, ch)
	}

	for {
		ch <- rand.IntN(10_000)
		time.Sleep(500 * time.Millisecond)
	}
}
