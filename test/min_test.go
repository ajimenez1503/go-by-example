package main

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	result := IntMin(2, -2)
	if result != -2 {
		t.Errorf("IntMin(2, -2) = %d, expected -2 ", result)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b     int
		expected int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("IntMin(%d,%d)", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			result := IntMin(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("IntMin(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
