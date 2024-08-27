// Simple Package for basic math opertaions
package math_ops

import (
	"fmt"
)
func Add(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (float64) {
	if b == 0 {
        fmt.Println("devision by zero is not allowed")
		return 0
	}
	return float64(a) / float64(b)
}


