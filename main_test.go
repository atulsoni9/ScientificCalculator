package main

import (
	"math"
	"testing"
)

// --- Basic Operations ---

func TestAdd(t *testing.T) {
	if Add(10, 5) != 15 {
		t.Error("Add failed")
	}
}
func TestSubtract(t *testing.T) {
	if Subtract(10, 5) != 5 {
		t.Error("Subtract failed")
	}
}
func TestMultiply(t *testing.T) {
	if Multiply(10, 5) != 50 {
		t.Error("Multiply failed")
	}
}

func TestDivide(t *testing.T) {
	res, err := Divide(10, 2)
	if err != nil || res != 5 {
		t.Error("Divide failed")
	}

	// Test Edge Case: Division by Zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Should have returned error for division by zero")
	}
}

// --- Scientific Operations ---

func TestPower(t *testing.T) {
	if Power(2, 3) != 8 {
		t.Error("Power failed")
	}
}

func TestSquareRoot(t *testing.T) {
	res, err := SquareRoot(25)
	if err != nil || res != 5 {
		t.Error("SquareRoot failed")
	}

	// Test Edge Case: Negative Square Root
	_, err = SquareRoot(-1)
	if err == nil {
		t.Error("Should have returned error for negative square root")
	}
}

func TestLogarithm(t *testing.T) {
	res, _ := Logarithm(math.E) // ln(e) = 1
	if res != 1 {
		t.Error("Logarithm failed")
	}

	// Test Edge Case: Log of 0
	_, err := Logarithm(0)
	if err == nil {
		t.Error("Should have returned error for ln(0)")
	}
}

func TestFactorial(t *testing.T) {
	if Factorial(0) != 1 {
		t.Error("Factorial(0) failed")
	}
	if Factorial(5) != 120 {
		t.Error("Factorial(5) failed")
	}
	//test
}
