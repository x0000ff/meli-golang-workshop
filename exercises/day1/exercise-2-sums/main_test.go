package main

import "testing"

func TestSum(t *testing.T) {

	tests := []struct {
		input    int
		expected int
	}{
		{-123, 0},
		{-1, 0},
		{1, 0},
		{3, 3},
		{5, 8},
		{6, 14},
		{9, 23},
		{10, 33},
	}

	for _, test := range tests {
		real := sum(test.input)
		if real != test.expected {
			t.Errorf("Expected %d, got %d for input %d", test.expected, real, test.input)
		}
	}
}
