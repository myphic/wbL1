package main

import "fmt"

/*
	Проблемы кода:
		1. Ненужная глобальная переменная
		2. Нет реализации createHugeString
		3. justString использует 100 элементов из 1024 (строка это слайс байтов доступный только для чтения)
*/

func createHugeString(size int) string {
	result := ""
	for i := 0; i < size; i++ {
		result += "123"
	}
	return result
}

func someFunc() {
	var justString string
	v := createHugeString(1 << 10)
	c := make([]byte, 100)
	copy(c, v[:100])
	justString = string(c)
	fmt.Print(justString)
}

func main() {
	someFunc()
}
