package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "overlap",
			a:    []int{1, 2, 3},
			b:    []int{2, 3, 4},
			want: []int{2, 3},
		},
		{
			name: "duplicates",
			a:    []int{2, 2, 3, 4},
			b:    []int{2, 2, 2, 4},
			want: []int{2, 4},
		},
		{
			name: "no intersection",
			a:    []int{1, 5},
			b:    []int{2, 4},
			want: nil,
		},
		{
			name: "empty input",
			a:    nil,
			b:    []int{1, 2, 3},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
