// Simple Package for basic math opertaions
package math_ops

// Add two numbers
func Add( a,b int ) int {
	return a + b 
}

// Substract two numbers
func Substract( a,b int ) int {
	return a - b
}

// Multiply two numbers
func Multiply(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) * b.(int)
	default:
		return 0
	}
}


func Divide( a,b int ) int {
	if b == 0 {
		return 0
	}
	return a / b
}
