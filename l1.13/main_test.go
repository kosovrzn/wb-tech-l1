package main

import "testing"

func TestSwapXOR(t *testing.T) {
	tests := []struct {
		name   string
		a, b   int
		wantA  int
		wantB  int
	}{
		{name: "positive", a: 5, b: 7, wantA: 7, wantB: 5},
		{name: "negative", a: -10, b: 3, wantA: 3, wantB: -10},
		{name: "zero_and_value", a: 0, b: 9, wantA: 9, wantB: 0},
		{name: "both_zero", a: 0, b: 0, wantA: 0, wantB: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := swapXOR(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Fatalf("swapXOR(%d, %d) = (%d, %d), want (%d, %d)", tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}
