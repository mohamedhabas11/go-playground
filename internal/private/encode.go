package private

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// EncodeData encodes the given data to a base64 string
func EncodeData(data interface{}) (string, error) {
	// Convert the data to a JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %w", err)
	}

	// Encode the JSON string to a base64 string
	encodedData := base64.StdEncoding.EncodeToString(jsonData)

	return encodedData, nil
}

// DecodeData decodes the base64 string to the original data structure
func DecodeData(encodedData string, data interface{}) error {
	// Decode the base64 string
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return fmt.Errorf("failed to decode data: %w", err)
	}

	// Unmarshal the decoded data into the provided data structure
	if err := json.Unmarshal(decodedData, data); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return nil
}

// Run runs a test on the EncodeData and DecodeData functions
func Run() {
	// Test data
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	// Encode the data
	encodedData, err := EncodeData(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}

	fmt.Println("Encoded data:", encodedData)

	// Decode the data
	var decodedData map[string]string
	err = DecodeData(encodedData, &decodedData)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}

	fmt.Println("Decoded data:", decodedData)
}
