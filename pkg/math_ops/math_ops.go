// Simple Package for basic math opertaions
package math_ops

import (
	"fmt"
)

// Add Any number of integers
func Add(numbers ...interface{}) float64 {
	var sum float64
	for _, num := range numbers {
		switch num := num.(type) {
		case int:
			sum += float64(num)
		case float64:
			sum += num
		default:
			fmt.Println("Invalid Type")
		}
	}
	return sum
}

// Subtract Any number of integers
func Subtract(numbers ...interface{}) float64 {
	var sum float64
	for _, num := range numbers {
		switch num := num.(type) {
		case int:
			sum -= float64(num)
		case float64:
			sum -= num
		default:
			fmt.Println("Invalid Type")
		}
	}
	return sum
}

// Multiply Any number of integers
func Multiply(numbers ...interface{}) float64 {
	var sum float64 = 1
	for _, num := range numbers {
		switch num := num.(type) {
		case int:
			sum *= float64(num)
		case float64:
			sum *= num
		default:
			fmt.Println("Invalid Type")
		}
	}
	return sum
}

// Divide Any number of integers
func Divide(numbers ...interface{}) float64 {
	// example: 2 / 2 / 2 = 0.5
	// example: 2 / 2 / 0.0 = infinity
	// example: 2 / 2 / 0 = infinity
	// example: 0 / 2 / 2 = 0
	var sum float64
	for i, num := range numbers {
		switch num := num.(type) {
		case int:
			if i == 0 {
				sum = float64(num)
			} else {
				sum /= float64(num)
			}
		case float64:
			if i == 0 {
				sum = num
			} else {
				sum /= num
			}
		default:
			fmt.Println("Invalid Type")
		}
	}
	return sum
}
