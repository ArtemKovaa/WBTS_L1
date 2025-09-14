package main

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// First way
	wg.Go(func() { conditionalExit() })

	// Second way
	stop := make(chan bool)
	wg.Go(func() { notifyChannelExit(stop) })
	time.Sleep(2 * time.Second)
	stop <- true

	// Third way
	ctx, cancel := context.WithCancel(context.Background())
	wg.Go(func() { contextExit(ctx) })
	time.Sleep(3 * time.Second)
	cancel()

	// Fourth way
	wg.Go(func() { runtimeGoExit() })

	// Fifth way
	wg.Go(func() { timeExit() })

	// Sixth way
	ch := make(chan int)
	wg.Go(func() { closeChannelExit(ch) })
	for i := range 5 {
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)

	// Wait all goroutines before last way
	wg.Wait()

	// Seventh way - deadlock with panic
	deadlock := make(chan string)
	deadlockExit(deadlock)
	log.Println("unreachable")
}

func conditionalExit() {
	for {
		v := rand.Intn(100)
		if v > 70 {
			log.Printf("Value %d more than 70, goroutine is stopping", v)
			return
		}
		log.Printf("Value %d is less than or equals to 70, goroutine keep working", v)
		time.Sleep(500 * time.Millisecond)
	}
}

func notifyChannelExit(stop <-chan bool) {
	for {
		select {
		case <-time.After(1 * time.Second):
			log.Println("Goroutine with notify channel is working")
		case <-stop:
			log.Println("Goroutine with notify channel is stopping")
			return
		}
	}
}

func contextExit(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			log.Println("Goroutine with context is working")
		case <-ctx.Done():
			log.Println("Goroutine with context is stopping")
			return
		}
	}
}

func runtimeGoExit() {
	log.Println("Goroutine with runtime.Goexit is working")
	time.Sleep(3 * time.Second)
	log.Println("Goroutine with runtime.Goexit is stopping")
	runtime.Goexit()
	log.Println("unreachable")
}

func timeExit() {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-timeout:
			log.Println("Goroutine with timeout is stopping")
			return
		default:
			log.Println("Goroutine with timeout is working")
			time.Sleep(1 * time.Second)
		}
	}
}

func closeChannelExit(ch <-chan int) {
	for v := range ch {
		log.Println("Goroutine with reading channel is working. Got value:", v)
	}
	log.Println("Goroutine with reading channel is stopping")
}

func deadlockExit(ch <-chan string) {
	for str := range ch {
		log.Println("Goroutine with deadlock got ", str)
	}
}
