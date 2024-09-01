package dict_utils

import "sort"

// Function to return dictionary keys
func GetKeys(dict map[string]string) []string {
	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	return keys
}

// Function to return dictionary keys in order
func GetKeysInOrder(dict map[string]string) []string {
	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// Function to return sorted dictionary object
func SortDict(dict map[string]string) map[string]string {
	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sorted_dict := make(map[string]string)
	for _, k := range keys {
		sorted_dict[k] = dict[k]
	}
	return sorted_dict
}

// Function that combines two dictionaries
func CombineDicts(dict1 map[string]string, dict2 map[string]string) map[string]string {
	combined_dict := make(map[string]string)
	for k, v := range dict1 {
		combined_dict[k] = v
	}
	for k, v := range dict2 {
		combined_dict[k] = v
	}
	return combined_dict
}
