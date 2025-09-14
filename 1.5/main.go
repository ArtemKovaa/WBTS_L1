package main

import (
	"log"
	"os"
	"strconv"
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

	ch := make(chan int)

	go producer(n, ch)

	for v := range ch {
		log.Printf("Consumed: %d", v)
	}
	
	log.Println("Program successfully executed")
}

func producer(workingTime int, ch chan<- int) {
	timeout := time.After(time.Duration(workingTime) * time.Second)
	for i := 0;;i++ {
		select {
		case <-timeout:
			log.Printf("Producer stopped after %d seconds", workingTime)
			close(ch)
			return
		default:
			ch <- i
			log.Printf("Produced: %d", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
