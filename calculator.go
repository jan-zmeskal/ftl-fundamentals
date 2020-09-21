// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes arbitrary amount o numbers returns the result of adding them together.
func Add(nums ...float64) (total float64) {
	for _, n := range nums {
		total += n
	}
	return
}

// Subtract takes arbitrary amount of numbers and returns the result of subtracting
// all of the from the first one.
func Subtract(nums ...float64) (total float64) {
	total = nums[0]
	for i := 1; i < len(nums); i++ {
		total -= nums[i]
	}
	return
}

// Multiply takes arbitrary amount of numbers and returns the result
// of multiplying the first one by every other.
func Multiply(nums ...float64) (total float64) {
	total = nums[0]
	for i := 1; i < len(nums); i++ {
		total *= nums[i]
	}
	return
}

// Divide takes arbitrary amount of numbers and returns the result
// of dividing the fist one by every other.
func Divide(nums ...float64) (total float64, err error) {
	total = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0, fmt.Errorf("%f/%f is invalid, cannot divide by zero", total, nums[i])
		}
		total /= nums[i]
	}
	return total, err
}

// Sqrt takes one number and returns its square root
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("%f is invalid input for Sqrt, it's negative number", a)
	}
	return math.Sqrt(a), nil
}
