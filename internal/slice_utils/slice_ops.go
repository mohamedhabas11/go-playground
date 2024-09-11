package slice_utils

// GetItem retrieves an item from a slice at a given position.
// If the index is out of range, it returns the zero value for the type of the slice elements.
func GetItem[T any](slice []T, index int) T {
	var zeroValue T
	if index < 0 || index >= len(slice) {
		return zeroValue
	}
	return slice[index]
}

// SetItem writes an item to a slice at a given position, overwriting an existing value.
// If the index is out of range, including negative indices, the value is appended to the end of the slice.
func SetItem[T any](slice []T, index int, value T) []T {
	if index < 0 || index >= len(slice) {
		return append(slice, value) // Append value if index is out of range
	}
	// If index is within range, replace the item
	return append(slice[:index], append([]T{value}, slice[index+1:]...)...)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems[T any](slice []T, values ...T) []T {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
// If the index is out of range, it returns the slice unchanged.
func RemoveItem[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice // No changes if index is out of range
	}
	return append(slice[:index], slice[index+1:]...)
}
