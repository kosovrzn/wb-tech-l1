package main

import (
	"fmt"
	"sync"
)

func squareNumbers(numbers []int) []int {
	results := make([]int, len(numbers))

	var wg sync.WaitGroup

	for i, n := range numbers {
		wg.Add(1)

		go func(idx, value int) {
			defer wg.Done()
			results[idx] = value * value
		}(i, n)
	}

	wg.Wait()

	return results
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	for _, square := range squareNumbers(numbers) {
		fmt.Println(square)
	}
}
