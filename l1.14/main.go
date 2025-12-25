package main

import (
	"fmt"
	"reflect"
)

func detectType(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}, <-chan interface{}, chan<- interface{}:
		return "chan"
	default:
		if value == nil {
			return "unknown"
		}

		if reflect.TypeOf(value).Kind() == reflect.Chan {
			return "chan"
		}

		return "unknown"
	}
}

func main() {
	samples := []interface{}{
		42,
		"hello",
		true,
		make(chan int),
		3.14,
	}

	for _, sample := range samples {
		fmt.Printf("Value: %-8v Type: %s\n", sample, detectType(sample))
	}
}
