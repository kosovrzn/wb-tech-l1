package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const maxBitIndex = 63

func main() {
	number := flag.Int64("n", 0, "source int64 value")
	bit := flag.Int("bit", 0, "bit index to update (0-63)")
	value := flag.Int("value", 1, "target bit value (0 or 1)")
	flag.Parse()

	result, err := setBit(*number, *bit, *value)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("input : %d (0b%b)\n", *number, *number)
	fmt.Printf("result: %d (0b%b)\n", result, result)
}

// setBit toggles the bit at the provided index of number and returns the updated value.
func setBit(number int64, index int, bitValue int) (int64, error) {
	if index < 0 || index > maxBitIndex {
		return 0, fmt.Errorf("bit index must be between 0 and %d", maxBitIndex)
	}
	if bitValue != 0 && bitValue != 1 {
		return 0, errors.New("bit value must be 0 or 1")
	}

	mask := int64(1) << index
	if bitValue == 1 {
		return number | mask, nil
	}
	return number &^ mask, nil
}
