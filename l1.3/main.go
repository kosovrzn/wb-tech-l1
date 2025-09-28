package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	workers := flag.Int("workers", 1, "number of worker goroutines")
	delay := flag.Duration("delay", 200*time.Millisecond, "pause between generated values")
	flag.Parse()

	if *workers <= 0 {
		fmt.Fprintln(os.Stderr, "workers must be greater than 0")
		os.Exit(1)
	}

	jobs := make(chan int)

	var wg sync.WaitGroup
	wg.Add(*workers)
	for i := 0; i < *workers; i++ {
		go func(id int) {
			defer wg.Done()
			for v := range jobs {
				fmt.Printf("worker %d processed value %d\n", id, v)
			}
		}(i + 1)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("started %d workers, press Ctrl+C to stop...\n", *workers)

	counter := 0
	for {
		select {
		case <-sigCh:
			close(jobs)
			wg.Wait()
			fmt.Println("shutdown complete")
			return
		case jobs <- counter:
			counter++
			time.Sleep(*delay)
		}
	}
}
