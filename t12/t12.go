package main

import "fmt"

func removeDuplicates(arr []string) []string {
	m := make(map[string]bool)
	result := make([]string, len(arr))
	for _, value := range arr {
		if _, ok := m[value]; !ok {
			m[value] = true
			result = append(result, value)
		}
	}
	return result
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "three"}
	fmt.Println(removeDuplicates(arr))
}
