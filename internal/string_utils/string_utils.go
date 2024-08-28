package string_utils

import (
	"fmt"
	"strings"
)

func GreetTheMan(name string) string {
	fmt.Println("Hello, " + name)
	return "Salamo 3alikom, " + name
}

func SubTesting(str string) {
	substring := "Go"
    contains := strings.Contains(str, substring)

	if contains {
		fmt.Println("Substring found!")
	} else {
		fmt.Println("Substring not found!")
	}

	count := strings.Count(str, substring)
	fmt.Println("Occurences: ", count)
}
