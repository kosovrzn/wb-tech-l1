package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Printf("Words: %v\n", words)

	unique := uniqueWords(words)

	ordered := make([]string, 0, len(unique))
	for w := range unique {
		ordered = append(ordered, w)
	}
	sort.Strings(ordered)

	fmt.Printf("Unique set: %v\n", ordered)
}

func uniqueWords(words []string) map[string]struct{} {
	result := make(map[string]struct{}, len(words))
	for _, w := range words {
		result[w] = struct{}{}
	}
	return result
}
