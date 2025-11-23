package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10

	groups := groupByStep(temperatures, step)

	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}

// groupByStep clusters values into buckets of width step; negative values round toward zero.
func groupByStep(values []float64, step int) map[int][]float64 {
	result := make(map[int][]float64)
	if step <= 0 {
		return result
	}

	stepFloat := float64(step)
	for _, v := range values {
		bucket := int(v/stepFloat) * step
		result[bucket] = append(result[bucket], v)
	}

	return result
}
