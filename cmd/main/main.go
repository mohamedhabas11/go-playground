package main

import (
	"fmt"
	// selfmade packages go here
	"github.com/mohamedhabas11/go-playground/pkg/math_ops"
)


func main() {
    fmt.Println("Hello, World!")
	
	// use the imported thing
	result := math_ops.Add(2, 3)
	fmt.Printf("2 + 3 = %d\n", result)

	float_result := math_ops.Multiply(2, 3)
	fmt.Printf("2 * 3 = %f\n", float_result)

	result = math_ops.Divide(2, 3)
	fmt.Printf("2 / 3 = %d\n", result)

	fmt.Print(result.(type))
}
