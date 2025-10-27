package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter wraps a map with a mutex so concurrent goroutines can update it.
type SafeCounter struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		data: make(map[string]int),
	}
}

func (c *SafeCounter) Add(key string, delta int) {
	c.mu.Lock()
	c.data[key] += delta
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *SafeCounter) Snapshot() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make(map[string]int, len(c.data))
	for k, v := range c.data {
		result[k] = v
	}
	return result
}

func main() {
	const (
		workers    = 8
		increments = 25_000
	)

	counter := NewSafeCounter()
	var wg sync.WaitGroup
	wg.Add(workers)

	stopReader := make(chan struct{})
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Printf("intermediate total=%d\n", counter.Value("total"))
			case <-stopReader:
				return
			}
		}
	}()

	for workerID := 0; workerID < workers; workerID++ {
		id := workerID
		go func() {
			defer wg.Done()
			for i := 0; i < increments; i++ {
				counter.Add("total", 1)
				counter.Add(fmt.Sprintf("worker-%d", id), 1)
			}
		}()
	}

	wg.Wait()
	close(stopReader)

	snapshot := counter.Snapshot()
	fmt.Printf("final total=%d (expected %d)\n", snapshot["total"], workers*increments)
	for i := 0; i < workers; i++ {
		key := fmt.Sprintf("worker-%d", i)
		fmt.Printf("%s=%d\n", key, snapshot[key])
	}
}
