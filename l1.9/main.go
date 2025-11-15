package main

import "fmt"

func main() {
	numbers := []int{3, 5, 7, 11, 13, 17}

	source := make(chan int)
	doubled := make(chan int)

	// Stage 1: feed numbers into the pipeline.
	go func() {
		defer close(source)
		for _, v := range numbers {
			source <- v
		}
	}()

	// Stage 2: read numbers from source, multiply them, and forward downstream.
	go func() {
		defer close(doubled)
		for v := range source {
			doubled <- v * 2
		}
	}()

	// Final stage: consume results and print them to stdout.
	for v := range doubled {
		fmt.Println(v)
	}
}
