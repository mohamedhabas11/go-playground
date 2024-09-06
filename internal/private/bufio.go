package private

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchAndReplace(searchStr, replaceStr, filePath string) error {
	// Open file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file for reading: %w", err)
	}
	defer file.Close()

	// Read the file line by line
	var result strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Perform search and replace
		updatedLine := strings.ReplaceAll(line, searchStr, replaceStr)
		result.WriteString(updatedLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Open file for writing (create or truncate)
	file, err = os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error opening file for writing: %w", err)
	}
	defer file.Close()

	// Write updated content to the file
	_, err = file.WriteString(result.String())
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}
	fmt.Println(result.String())
	return nil
}
