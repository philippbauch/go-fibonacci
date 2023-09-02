package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        uint
		expected uint64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{5, 5},
		{10, 55},
		{20, 6765},
		{30, 832040},
		{40, 102334155},
		{50, 12586269025},
	}

	for _, test := range tests {
		actual, err := Fibonacci(test.n)
		if err != nil {
			t.Errorf("Fibonacci(%d): %s", test.n, err)
		}
		if actual != test.expected {
			t.Errorf("Fibonacci(%d): expected %d, actual %d", test.n, test.expected, actual)
		}
	}
}
