package private

import "fmt"

func Fib() func() int { // function taking a function as input
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func Closure() {
	_closure := make([]func(), 4)

	for _index := 0; _index < 4; _index++ {
        // Closure capture: assign value to local variable to obtain new memory address
		_index_closure := _index
        // Pointer Printing: The closure prints the value of _index_closure and its memory address.
		_closure[_index] = func() {
			fmt.Printf("%d @ %p\n", _index_closure, &_index_closure)
		}
	}

	for _index := 0; _index < 4; _index++ {
		_closure[_index]()
	}

}
