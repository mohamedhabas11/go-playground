package main

import (
	// built-in packages go here
	"fmt"

	// selfmade packages go here
	"github.com/mohamedhabas11/go-playground/internal/dict_utils"
	"github.com/mohamedhabas11/go-playground/internal/string_utils"

	// Global packages go here
	"github.com/mohamedhabas11/go-playground/pkg/math_ops"
	"github.com/mohamedhabas11/go-playground/pkg/net_ops"
)

func main() {
	fmt.Println("Hello, World!")

	// use the imported internal package
	name := "Mohamed"
	greet := string_utils.GreetTheMan(name)
	fmt.Println(greet)

	// use the imported global package
	result := math_ops.Add(2, 3)
	fmt.Printf("2 + 3 = %f\n", result)

	result = math_ops.Subtract(2, 3, 4)
	fmt.Printf("2 - 3 - 4 = %f\n", result)

	result = math_ops.Multiply(5, 5, 10)
	fmt.Printf("5 * 5 * 10 = %f\n", result)

	result = math_ops.Divide(2, 2, 2)
	fmt.Printf("2 / 2 / 2 = %f\n", result)

	// use the imported internal package
	test_dict := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	// prints [key2 key3 key1] out of order ?
	fmt.Println(dict_utils.GetKeys(test_dict))
	// This is because the map is unordered
	// To get the keys in order, you can sort them
	// before returning them

	// prints [key1 key2 key3] in order
	fmt.Println(dict_utils.GetKeysInOrder(test_dict))

	// prints map[key1:value1 key2:value2 key3:value3] in order
	fmt.Println(dict_utils.SortDict(test_dict))

	// combine two dictionaries
	test_dict2 := map[string]string{
		"key4": "value4",
		"key5": "value5",
		"key6": "value6",
	}

	fmt.Println(dict_utils.CombineDicts(test_dict, test_dict2))

	iface := "bridge100"
	ipaddress, err := net_ops.CheckInterface(iface)
	if err != nil {
		fmt.Printf("Error checking interface: %v\n", err)
		return
	}
	fmt.Printf("interface: %v  ipaddr:%v\n", iface, ipaddress)

}
