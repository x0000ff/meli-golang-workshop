package app

// package app__test

import (
	"testing"
)

func TestGetHelp(t *testing.T) {

	expected := `Usage:
$ go run my-app.go <MAX>
$ go run my-app.go 31

Only integer values allowed`

	actual := GetHelp("my-app.go")
	if actual != expected {
		t.Fatalf("Expected: %q, got: %q", expected, actual)
	}
}

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
		actual := sum(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d for input %d", test.expected, actual, test.input)
		}
	}
}
