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
	fmt.Printf("2 + 3 = %d\n", result)

}
