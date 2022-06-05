package main

import (
	"fmt"
	"sort"
)

func BinarySearch(haystack []int, needle int) (success bool) {
	sort.Ints(haystack)
	lowKey := 0
	highKey := len(haystack) - 1
	if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
		return
	}
	for lowKey <= highKey {
		mid := (lowKey + highKey) / 2
		if haystack[mid] == needle {
			success = true
			return
		}
		if haystack[mid] < needle {
			lowKey = mid + 1
			continue
		}
		if haystack[mid] > needle {
			highKey = mid - 1
		}
	}
	return
}

func main() {
	fmt.Println(BinarySearch([]int{100, 2, 3, -100, 5, 10, -132}, -100))
}
