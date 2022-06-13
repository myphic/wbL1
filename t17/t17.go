package main

import (
	"fmt"
	"sort"
)

func BinarySearch(haystack []int, needle int) (success bool) {
	//Бинарный поиск требует отсортированности массива
	sort.Ints(haystack)
	lowKey := 0
	highKey := len(haystack) - 1
	//Значение не в диапазоне
	if (haystack[lowKey] > needle) || (haystack[highKey] < needle) {
		return
	}
	for lowKey <= highKey {
		mid := (lowKey + highKey) / 2 //устанавливаем середину
		if haystack[mid] == needle {
			success = true //значение найдено
			return
		}
		if haystack[mid] < needle {
			//если поиск больше центра выбираем правый подмассив
			lowKey = mid + 1
			continue
		}
		if haystack[mid] > needle {
			//если поиск меньше центра выбираем левый подмассив
			highKey = mid - 1
		}
	}
	return
}

func main() {
	fmt.Println(BinarySearch([]int{100, 2, 3, -100, 5, 10, -132}, -100))
}
