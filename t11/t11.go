package main

import "fmt"

func intersect(first, second []int) []int {
	m := make(map[int]int)
	var result []int

	for _, value := range first {
		if _, ok := m[value]; !ok {
			m[value] = 1
		} else {
			m[value] += 1
		}
	}
	for _, value := range second {
		if count, ok := m[value]; ok && count > 0 {
			m[value] -= 1
			result = append(result, value)
		}
	}
	return result
}

func main() {
	first := []int{3, 4, 5, 220, -4, 3, 0}
	second := []int{-30, 4, 5, -20}
	fmt.Println(intersect(first, second))
}
