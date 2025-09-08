package main

import (
	"context"
	"log"
	"math/rand/v2"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	wg := sync.WaitGroup{}
	
	wg.Go(func() {
		worker_1(ctx)
	})
	wg.Go(func() {
		worker_2(ctx)
	})

	wg.Wait()
}

func worker_1(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		log.Println("Worker-1: Application is running, waiting for SIGINT...")
	case <-ctx.Done():
		log.Println("Worker-1: Graceful shutdown started...")

		log.Println("Worker-1: Processing remaining requests...")
		time.Sleep(3 * time.Second)
		log.Println("Worker-1: Successfully processed all reuqests")

		log.Println("Worker-1: Stopping DB connections...")
		time.Sleep(1 * time.Second)
		log.Println("Worker-1: DB connections stopped successfully...")

		log.Println("Worker-1: Graceful shutdown finished successfully")

		return
	}
}

func worker_2(ctx context.Context) {
	for {
		select {
		case <-time.After(100 * time.Millisecond):
			n := rand.IntN(10_000)
			log.Println("Worker-2: Doing some job with id", n)
			time.Sleep(3 * time.Second)
		case <-ctx.Done():
			log.Println("Worker-2: Shutdown gracefully")
			return
		}
	}
}
