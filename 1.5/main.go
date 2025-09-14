package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Enter program execution time as first argument")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Incorrect program execution time")
	}

	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Go(func() {
		consumer(n, ch)
	})
	wg.Go(func() {
		producer(n, ch)
	})

	wg.Wait()
	log.Println("Program successfully executed")
}

func producer(workingTime int, ch chan<- int) {
	timeout := time.After(time.Duration(workingTime) * time.Second)
	for i := 0;;i++ {
		select {
		case <-timeout:
			log.Printf("Producer stopped after %d seconds", workingTime)
			return
		default:
			ch <- i
			log.Printf("Produced: %d", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func consumer(workingTime int, ch <-chan int) {
	timeout := time.After(time.Duration(workingTime) * time.Second)
	for {
		select {
		case v := <-ch:
			log.Printf("Consumed: %d", v)
		case <-timeout:
			log.Printf("Consumer stopped after %d seconds", workingTime)
			return
		}
	}
}
