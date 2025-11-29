package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5, 5, 10, 15}
	b := []int{2, 3, 4, 5, 5, 8, 10}

	fmt.Printf("A = %v\n", a)
	fmt.Printf("B = %v\n", b)

	fmt.Printf("Intersection = %v\n", intersection(a, b))
}

func intersection(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	seen := make(map[int]struct{}, len(b))
	for _, v := range b {
		seen[v] = struct{}{}
	}

	var result []int
	for _, v := range a {
		if _, ok := seen[v]; ok {
			result = append(result, v)
			delete(seen, v)
		}
	}

	return result
}
