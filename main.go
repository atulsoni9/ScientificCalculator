package main

import (
	"fmt"
	"math"
	"os"
)

func Add(a, b float64) float64      { return a + b }
func Subtract(a, b float64) float64 { return a - b }
func Multiply(a, b float64) float64 { return a * b }
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero not allowed")
	}
	return a / b, nil
}
func Power(a, b float64) float64 { return math.Pow(a, b) }
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("negative number")
	}
	return math.Sqrt(a), nil
}
func Logarithm(a float64) (float64, error) {
	if a <= 0 {
		return 0, fmt.Errorf("undefined for non-positive numbers")
	}
	return math.Log(a), nil
}
func Factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	for {
		var op int
		var num1, num2 float64

		fmt.Println("\n--- Scientific Calculator (CLI) ---")
		fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Divide")
		fmt.Println("5. Power (x^b)\n6. Square Root\n7. Logarithm (ln)\n8. Factorial\n9. Exit")
		fmt.Print("Choose an option (1-9): ")
		fmt.Scan(&op)

		if op == 9 {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		if op >= 1 && op <= 5 {
			fmt.Print("Enter first number: ")
			fmt.Scan(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scan(&num2)
		} else if op >= 6 && op <= 8 {
			fmt.Print("Enter number: ")
			fmt.Scan(&num1)
		}

		switch op {
		case 1:
			fmt.Printf("Result: %.4f\n", Add(num1, num2))
		case 2:
			fmt.Printf("Result: %.4f\n", Subtract(num1, num2))
		case 3:
			fmt.Printf("Result: %.4f\n", Multiply(num1, num2))
		case 4:
			res, err := Divide(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %.4f\n", res)
			}
		case 5:
			fmt.Printf("Result: %.4f\n", Power(num1, num2))
		case 6:
			res, err := SquareRoot(num1)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %.4f\n", res)
			}
		case 7:
			res, err := Logarithm(num1)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %.4f\n", res)
			}
		case 8:
			if num1 < 0 || num1 != math.Trunc(num1) {
				fmt.Println("Error: Factorial requires a non-negative integer")
			} else {
				fmt.Printf("Result: %d\n", Factorial(uint64(num1)))
			}
		default:
			fmt.Println("Invalid Option. Please try again.")
		}
	}
}
