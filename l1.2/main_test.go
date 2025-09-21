package main

import "testing"

func TestSquareNumbers(t *testing.T) {
	input := []int{2, 4, 6, 8, 10}
	want := []int{4, 16, 36, 64, 100}

	got := squareNumbers(input)

	if len(got) != len(want) {
		t.Fatalf("unexpected result length: got %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("unexpected value at index %d: got %d, want %d", i, got[i], want[i])
		}
	}
}
