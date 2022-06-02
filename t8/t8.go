package main

import (
	"fmt"
	"strconv"
)

func setBit(position int, number int64, setUnit bool) string {
	if setUnit {
		fmt.Printf("Число до преобразования: %s \n", strconv.FormatInt(number, 2))
		number = number | (1 << position)
		return strconv.FormatInt(number, 2)
	} else {
		fmt.Printf("Число до преобразования: %s \n", strconv.FormatInt(number, 2))
		number = number &^ (1 << position)
		return strconv.FormatInt(number, 2)
	}
}

func main() {
	var num int64 = 14
	fmt.Printf("Число после преобразования: %s", setBit(0, num, true))
}
