package main

import (
	"fmt"
	"math"
	"os"
)

// Factorial function (recursive)
func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
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

		// Scientific options 6, 7, and 8 only require one input
		if op >= 1 && op <= 5 {
			fmt.Print("Enter first number: ")
			fmt.Scan(&num1)
			if op != 6 && op != 7 && op != 8 {
				fmt.Print("Enter second number: ")
				fmt.Scan(&num2)
			}
		} else if op >= 6 && op <= 8 {
			fmt.Print("Enter number: ")
			fmt.Scan(&num1)
		}

		switch op {
		case 1:
			fmt.Printf("Result: %.4f\n", num1+num2)
		case 2:
			fmt.Printf("Result: %.4f\n", num1-num2)
		case 3:
			fmt.Printf("Result: %.4f\n", num1*num2)
		case 4:
			if num2 == 0 {
				fmt.Println("Error: Division by zero")
			} else {
				fmt.Printf("Result: %.4f\n", num1/num2)
			}
		case 5:
			fmt.Printf("Result: %.4f\n", math.Pow(num1, num2))
		case 6:
			if num1 < 0 {
				fmt.Println("Error: Cannot take root of negative number")
			} else {
				fmt.Printf("Result: %.4f\n", math.Sqrt(num1))
			}
		case 7:
			if num1 <= 0 {
				fmt.Println("Error: Logarithm undefined for non-positive numbers")
			} else {
				fmt.Printf("Result: %.4f\n", math.Log(num1))
			}
		case 8:
			if num1 < 0 || num1 != math.Trunc(num1) {
				fmt.Println("Error: Factorial requires a non-negative integer")
			} else {
				fmt.Printf("Result: %d\n", factorial(uint64(num1)))
			}
		default:
			fmt.Println("Invalid Option. Please try again.")
		}
	}
}
