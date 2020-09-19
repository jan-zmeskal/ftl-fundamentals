package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b, want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	var testCases = []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
		{a: -3, b: 5, want: 2},
		{a: -3, b: -4, want: -7},
		{a: 0, b: 0, want: 0},
		{a: 3.14, b: 2, want: 5.14}, // TODO: This fails
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 6
	got := calculator.Multiply(2, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
