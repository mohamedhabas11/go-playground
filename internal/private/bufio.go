package private

import (
	"fmt"
	"os"
	"strings"
)

// OpenFile reads the entire file content and returns it as a string
func OpenFile(filePath string) (string, error) {
	// Read the file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(content), nil
}

// SearchAndReplace performs search and replace operations on a file and writes the result back
func SearchAndReplace(searchStr, replaceStr, filePath string) error {
	// Read the entire file content using OpenFile
	content, err := OpenFile(filePath)
	if err != nil {
		return err
	}

	// Perform search and replace
	updatedContent := strings.ReplaceAll(content, searchStr, replaceStr)

	// If no changes were made, return early
	if updatedContent == content {
		return nil
	}

	// Write updated content to a temporary file
	tmpFilePath := filePath + ".tmp"
	err = os.WriteFile(tmpFilePath, []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing to temporary file: %w", err)
	}

	// Replace the original file with the temporary file
	err = os.Rename(tmpFilePath, filePath)
	if err != nil {
		return fmt.Errorf("error replacing original file: %w", err)
	}

	fmt.Println("File successfully updated.")
	return nil
}

// WordCount counts the number of words in a file
func WordCount(filePath string) (int, error) {
	// Read the entire file content using OpenFile
	content, err := OpenFile(filePath)
	if err != nil {
		return 0, err
	}

	// Count the words in the file content
	words := strings.Fields(content)
	return len(words), nil
}

// UniqueWordCount counts the number of unique words in a file
func UniqueWordCount(filePath string) (int, error) {
	// Read the entire file content using OpenFile
	content, err := OpenFile(filePath)
	if err != nil {
		return 0, err
	}

	// Use a map to store unique words, using the word as the 'key' of the map, and an empty struct as the 'value'
	// This allows us to easily check for uniqueness, and appearantly the empty struct takes up no memory, and this is quite common in Go !!
	uniqueWords := make(map[string]struct{})
	words := strings.Fields(strings.ToLower(content))
	for _, word := range words {
		uniqueWords[word] = struct{}{}
	}

	return len(uniqueWords), nil
}
