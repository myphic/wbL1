package main

import "fmt"

//Функция для создания множества с помощью мапы
func removeDuplicates(arr []string) []string {
	m := make(map[string]bool)
	result := make([]string, len(arr))
	for index, value := range arr {
		if _, ok := m[value]; !ok {
			m[value] = true
			result[index] = value
		}
	}
	return result
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "three"}
	fmt.Println(removeDuplicates(arr))
}
