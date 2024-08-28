// Simple Package for basic math opertaions
package math_ops

import (
	"fmt"
	"math"
)

// Add two numbers
func Add( a, b int ) int {
	return a + b 
}

// Substract two numbers
func Substract( a, b int ) int {
	return a - b
}

func Multiply( a, b float64 ) float64 {
	return a * b
}

func Divide( a, b int ) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func PrintResult(result float64) {
	if math.Mod(result, 1) == 0 {
		fmt.Println(int(result))
	} else {
		fmt.Println(result)
	}
}
