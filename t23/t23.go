package main

import "fmt"

func removeSliceElem(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	arr := []int{1, 3, -40, 23, 41, 22, 41, 44, -2}
	newArr := removeSliceElem(arr, len(arr)-1)
	fmt.Println(newArr)
}
