// simple tricky string function
// Notes:
// 1. The package name should be the same as the folder name
// 2. The package name should be the same as the last part of the import path
// 3. Use Capitalized function names to export them to other packages like 'main'

package string_utils

import (
	"fmt"
	"strings"
)

// impelement GreetTheMan with printf
func GreetTheMan(name string) string {
	// return formatted string
	return fmt.Sprintf("Hello, %s!", name)
}

func SubTesting() {
	str := "Say, What?"
	substring := "Code & Go"

	contains := strings.Contains(str, substring)
	fmt.Println(contains)
}
