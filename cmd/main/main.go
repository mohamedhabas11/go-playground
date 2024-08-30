package main

import (
	// built-in packages go here
	"fmt"
	// selfmade packages go here
	"github.com/mohamedhabas11/go-playground/internal/string_utils"
	// Global packages go here
	"github.com/mohamedhabas11/go-playground/pkg/math_ops"
)

func main() {
	fmt.Println("Hello, World!")

	// use the imported internal package
	name := "Mohamed"
	greet := string_utils.GreetTheMan(name)
	fmt.Println(greet)

	// use the imported global package
	result := math_ops.Add(2, 3)
	fmt.Printf("2 + 3 = %f\n", result)

	result = math_ops.Subtract(2, 3, 4)
	fmt.Printf("2 - 3 - 4 = %f\n", result)

	result = math_ops.Multiply(5, 5, 10)
	fmt.Printf("5 * 5 * 10 = %f\n", result)

	result = math_ops.Divide(2, 2, 2)
	fmt.Printf("2 / 2 / 2 = %f\n", result)

}
