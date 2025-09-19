package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg := sync.WaitGroup{}

	wg.Go(func() { generator(nums, ch1) })
	wg.Go(func() { multiplier(ch1, ch2) })
	wg.Go(func() { printer(ch2) })

	wg.Wait()

	fmt.Println("Main goroutine is stopped")
}

func generator(nums []int, ch chan<- int) {
	for _, num := range nums {
		ch <- num
	}
	close(ch)
	fmt.Println("Generator is stopped")
}

func multiplier(ch1 <-chan int, ch2 chan<- int) {
	for val := range ch1 {
		ch2 <- val * 2
	}
	close(ch2)
	fmt.Println("Multiplier is stopped")
}


func printer(ch <-chan int) {
	for val := range ch {
		log.Printf("Got value: %d", val)
	}
	fmt.Println("Printer is stopped")
}
