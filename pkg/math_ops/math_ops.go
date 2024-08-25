// Simple Package for basic math opertaions
package math_ops

import "errors"

func Add(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return float64(a) / float64(b), nil
}
