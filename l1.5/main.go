package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	duration := flag.Duration("duration", 5*time.Second, "how long to run before stopping")
	interval := flag.Duration("interval", 500*time.Millisecond, "pause between generated values")
	flag.Parse()

	if *duration <= 0 {
		fmt.Fprintln(os.Stderr, "duration must be positive")
		os.Exit(1)
	}

	values := make(chan int)
	done := make(chan struct{})

	stop := time.After(*duration)
	go func() {
		<-stop
		close(done)
	}()

	go func() {
		defer close(values)

		ticker := time.NewTicker(*interval)
		defer ticker.Stop()

		next := 1
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				select {
				case values <- next:
					next++
				case <-done:
					return
				}
			}
		}
	}()

	for v := range values {
		fmt.Println(v)
	}

	fmt.Println("finished")
}
