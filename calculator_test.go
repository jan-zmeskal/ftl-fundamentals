package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
	"time"
)

const randIter int = 100

type testCase struct {
	a, b, want float64
	desc       string
}

type advTestCase struct {
	a, b, want  float64
	desc        string
	errExpected bool
}

type randomTestCase struct {
	a, b float64
}

func compare(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

func TestAdd(t *testing.T) {
	t.Parallel()
	var testCases = []testCase{
		{a: 2, b: 2, want: 4, desc: "Add two positive to get positive"},
		{a: 5, b: 0, want: 5, desc: "Add positive and zero to get positive"},
		{a: -3, b: 5, want: 2, desc: "Add negative and positive to get bigger"},
		{a: -3, b: -4, want: -7, desc: "Add two negative to get smaller"},
		{a: 0, b: 0, want: 0, desc: "Add two zeros to get zero"},
		{a: 3.14, b: 2, want: 5.14, desc: "Add one positive and fraction to get bigger fraction"},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if !compare(tc.want, got) {
			t.Errorf("%q failed. Add(%f, %f): want %f, got %f", tc.desc, tc.a, tc.b, tc.want, got)
		}
	}
}

// TODO: Rewrite using test cases
func TestSubtract(t *testing.T) {
	t.Parallel()
	var testCases = []testCase{
		{a: 2, b: 2, want: 0, desc: "Subtract positive from positive to get zero"},
		{a: 5, b: 0, want: 5, desc: "Subtract zero from positive to get the original positive"},
		{a: -3, b: 5, want: -8, desc: "Subtract positive from negative to get smaller"},
		{a: -3, b: -4, want: 1, desc: "Subtract negative from negative to get bigger"},
		{a: 0, b: 0, want: 0, desc: "Subtract zero from zero to get zero"},
		{a: 3.14, b: 2, want: 1.14, desc: "Subtract positive from fraction to get smaller fraction"},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if !compare(tc.want, got) {
			t.Errorf("%q failed. Subtract(%f, %f): want %f, got %f", tc.desc, tc.a, tc.b, tc.want, got)
		}
	}
}

// TODO: Rewrite using test cases
func TestMultiply(t *testing.T) {
	t.Parallel()
	var testCases = []testCase{
		{a: 2, b: 2, want: 4, desc: "Multiply two positive to get bigger positive"},
		{a: 5, b: 0, want: 0, desc: "Multiply positive and zero to get zero"},
		{a: -3, b: 5, want: -15, desc: "Multiply negative and positive to get bigger negative"},
		{a: -3, b: -4, want: 12, desc: "Multiply two negatives to get positive"},
		{a: 0, b: 0, want: 0, desc: "Multiply two zeros to get zero"},
		{a: 3.14, b: 2, want: 6.28, desc: "Multiply fraction and positive to get bigger fraction"},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if !compare(tc.want, got) {
			t.Errorf("%q failed. Multiply(%f, %f): want %f, got %f", tc.desc, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []advTestCase{
		{a: 2, b: 2, want: 1, desc: "Divide positive by positive to get positive", errExpected: false},
		{a: 5, b: 0, want: 0, desc: "Divide positive by zero to get error", errExpected: true},
		{a: 0, b: 5, want: 0, desc: "Divide zero by positive to get zero", errExpected: false},
		{a: -3, b: 5, want: -0.6, desc: "Divide negative by positive to get negative", errExpected: false},
		{a: -3, b: -4, want: 0.75, desc: "Divide two negative to get positive", errExpected: false},
		{a: 0, b: 0, want: 0, desc: "Divide zero by zero to get error", errExpected: true},
		{a: 3.14, b: 2, want: 1.57, desc: "Divide fraction by positive to get positive", errExpected: false},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("%q failed. Divide(%f, %f): unexpected error status: %v", tc.desc, tc.a, tc.b, errReceived)
		}
		if !tc.errExpected && !compare(tc.want, got) {
			t.Errorf("%q failed. Divide(%f, %f): want %f, got %f", tc.desc, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	var i int
	rand.Seed(time.Now().UnixNano())
	for i < randIter {
		tc := randomTestCase{a: rand.Float64() * 10, b: rand.Float64() * 20}
		want := tc.a + tc.b
		got := calculator.Add(tc.a, tc.b)
		if !compare(want, got) {
			t.Errorf("random test failed. Add(%f, %f): want %f, got %f", tc.a, tc.b, want, got)
		}
		i++
	}
}
