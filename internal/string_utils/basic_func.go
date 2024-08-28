// simple tricky string function

package basic_func

import (
	"fmt"
	"strings"
)

func greetTheMan(name string) string {
	fmt.Println("Hello, " + name)
	return "Hello, " + name
}

func subTesting() {
    str := "Say, What?"
    substring := "Code & Go"

    contains := strings.Contains(str, substring)
    fmt.Println(contains)
}
