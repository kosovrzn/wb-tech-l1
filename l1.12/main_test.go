package main

import (
	"reflect"
	"testing"
)

func TestUniqueWords(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  map[string]struct{}
	}{
		{
			name:  "deduplicates values",
			input: []string{"cat", "cat", "dog", "cat", "tree"},
			want: map[string]struct{}{
				"cat":  {},
				"dog":  {},
				"tree": {},
			},
		},
		{
			name:  "empty input",
			input: nil,
			want:  map[string]struct{}{},
		},
		{
			name:  "preserves case sensitivity",
			input: []string{"Cat", "cat"},
			want: map[string]struct{}{
				"Cat": {},
				"cat": {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueWords(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("uniqueWords(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
