package private

import (
	"fmt"
	"io"
	"os"
)

// ReadFile reads the contents of a file using the os package.
func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the contents of the file
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(data), nil
}

// AppendToFile appends data to an existing file.
func AppendToFile(path string, data string) error {
	// Open the file in append mode, creating the file if it doesn't exist.
	// Use the correct flags to ensure the file is opened in append mode.
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Append data to the file
	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to append to file: %w", err)
	}

	return nil
}
