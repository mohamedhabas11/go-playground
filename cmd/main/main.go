package main

import (
	"fmt"
	// selfmade packages go here
	"github.com/mohamedhabas11/go-playground/pkg/math_ops"
	"github.com/mohamedhabas11/go-playground/internal/string_utils"
)

func main() {
	string_utils.GreetTheMan("Mohamed")

	// use the imported thing
	result := math_ops.Add(2, 3)
	fmt.Printf("2 + 3 = %d\n", result)

	// use the other imported thing
	string_utils.SubTesting("Code & Go and then Go, because Go is awesome!")
}
