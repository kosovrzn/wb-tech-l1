package main

import "fmt"

func main() {
	a, b := 42, -7

	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)

	a, b = swapXOR(a, b)

	fmt.Printf("After swap:  a=%d, b=%d\n", a, b)
}

func swapXOR(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}
