package main

import "testing"

func TestDetectType(t *testing.T) {
	var recvOnly <-chan bool = make(chan bool)
	var sendOnly chan<- string = make(chan string)

	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{name: "int", value: 7, want: "int"},
		{name: "string", value: "go", want: "string"},
		{name: "bool", value: false, want: "bool"},
		{name: "chan int", value: make(chan int), want: "chan"},
		{name: "recv-only chan", value: recvOnly, want: "chan"},
		{name: "send-only chan", value: sendOnly, want: "chan"},
		{name: "unknown type", value: 3.14, want: "unknown"},
		{name: "nil", value: nil, want: "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectType(tt.value); got != tt.want {
				t.Fatalf("detectType(%v) = %s, want %s", tt.value, got, tt.want)
			}
		})
	}
}
