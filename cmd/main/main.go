package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	// Self-made packages
	"github.com/mohamedhabas11/go-playground/internal/dict_utils"
	"github.com/mohamedhabas11/go-playground/internal/private"
	"github.com/mohamedhabas11/go-playground/internal/string_utils"

	// Global packages
	"github.com/mohamedhabas11/go-playground/pkg/math_ops"
	"github.com/mohamedhabas11/go-playground/pkg/net_ops"
)

// echoMsg returns the message repeated a specified number of times.
// If no number is specified, it repeats the message twice.
func echoMsg(msg string, times ...int) string {
	repeatCount := 2 // Default repetition count
	if len(times) > 0 {
		repeatCount = times[0]
	}

	// Create a slice to hold the repeated messages
	repeatedMsgs := make([]string, repeatCount)
	for i := 0; i < repeatCount; i++ {
		repeatedMsgs[i] = msg
	}

	// Join the repeated messages with a space
	return strings.Join(repeatedMsgs, " ")
}

func main() {
	fmt.Println("Hello, World!")

	// Using the internal string_utils package
	name := "Mohamed"
	greet := string_utils.GreetTheMan(name)
	fmt.Println(greet)

	// Using the global math_ops package
	fmt.Printf("2 + 3 = %f\n", math_ops.Add(2, 3))
	fmt.Printf("2 - 3 - 4 = %f\n", math_ops.Subtract(2, 3, 4))
	fmt.Printf("5 * 5 * 10 = %f\n", math_ops.Multiply(5, 5, 10))
	fmt.Printf("2 / 2 / 2 = %f\n", math_ops.Divide(2, 2, 2))

	// Using the internal dict_utils package
	testDict := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	// Print map keys (unordered and ordered)
	fmt.Println("Unordered keys:", dict_utils.GetKeys(testDict))
	fmt.Println("Ordered keys:", dict_utils.GetKeysInOrder(testDict))
	fmt.Println("Sorted dictionary:", dict_utils.SortDict(testDict))

	// Combine two dictionaries
	testDict2 := map[string]string{
		"key4": "value4",
		"key5": "value5",
		"key6": "value6",
	}
	combinedDict := dict_utils.CombineDicts(testDict, testDict2)
	fmt.Println("Combined dictionaries:", combinedDict)

	// Using the global net_ops package
	iface := "bridge100"
	ipAddress, err := net_ops.CheckInterface(iface)
	if err != nil {
		log.Printf("Error checking interface: %v\n", err)
	} else {
		fmt.Printf("Interface: %v, IP Address: %v\n", iface, ipAddress)
	}

	// // Local echoMsg function usage
	fmt.Println(echoMsg("keep it simple!"))

	// // Using the internal private package
	// private.Run()

	// File operations using the private package
	filePath := "test_dir/test.txt"
	appendedData := "This is appended text!"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if err := private.AppendToFile(filePath, fmt.Sprintf("%v: %v\n", timestamp, appendedData)); err != nil {
		log.Fatalf("Error appending to file: %v", err)
	}
	// Read and print the contents of the file
	content, err := private.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	} else {
		fmt.Println("File contents:")
		fmt.Println(content)
	}

	domain_name := "github.com"
	ip, err := net_ops.DNSLookup(domain_name)
	if err != nil {
		fmt.Printf("Error resolving: %v", err)
	}
	fmt.Printf("%v resolves to %v\n", domain_name, ip)

	records, err := net_ops.DNSLookupNS(domain_name)
	if err != nil {
		fmt.Printf("Error resolving: %v", err)
		return
	}
	fmt.Println("NameServers:")
	for _, ns := range records {
		fmt.Println(ns)
	}

}
