package main

import (
	"log"
	"sync"
	"sync/atomic"
)

type Conter interface {
	Increment()
	Get() int64
}

type CounterMtx struct {
	value int64
	mtx   sync.RWMutex
}

func (c *CounterMtx) Increment() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.value++
}

func (c *CounterMtx) Get() int64 {
	c.mtx.RLock()
	defer c.mtx.RLock()

	return c.value
}

type CounterAtomic struct {
	value atomic.Int64
}

func (c *CounterAtomic) Increment() {
	c.value.Add(1)
}

func (c *CounterAtomic) Get() int64 {
	return c.value.Load()
}

func work(counter Conter) {
	for range 1_000_000 {
		counter.Increment()
	}
}

func main() {
	wg := sync.WaitGroup{}
	
	// First implementation
	counterMtx := CounterMtx{}
	for range 10 {
		wg.Go(func() { work(&counterMtx) })
	}
	wg.Wait()
	
	log.Println("Value for counter with mutex:", counterMtx.Get())

	// Second implementation
	counterAtomic := CounterAtomic{}
	for range 10 {
		wg.Go(func() { work(&counterAtomic) })
	}
	wg.Wait()
	log.Println("Value for counter with atomic:", counterAtomic.Get())
}
