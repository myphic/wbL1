package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(arr []int) []int {
	//Рекурсия не применяется к массиву в котором 1 или отсутствуют элементы
	if len(arr) < 2 {
		return arr
	}
	//левая и правая граница
	left, right := 0, len(arr)-1
	//случайный выбор опорного элемента
	rand.Seed(time.Now().UnixNano())
	pivot := rand.Int() % len(arr)
	//Передвигаем опорный и правый элемент
	arr[pivot], arr[right] = arr[right], arr[pivot]
	//Разбиение: перераспределение элементов в массиве таким образом, что элементы, меньшие опорного, помещаются перед ним, а большие или равные - после.
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	//Рекурсивно отсортировать массивы до опорного и после
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{-100, 2, 50, 0, -123, -1000, 102}
	QuickSort(arr)
	fmt.Println(arr)
}
