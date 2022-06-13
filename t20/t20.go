package main

import (
	"fmt"
	"strings"
)

func Reverse(input []string) []string {
	inputLen := len(input)
	output := make([]string, inputLen)

	for i, n := range input {
		j := inputLen - i - 1
		output[j] = n
	}

	return output
}

func wordReverse(s string) string {
	//Делим строку по пробелам
	splitStr := strings.Split(s, " ")
	//С помощью функции Reverse меняем элементы в нужном порядке и джойним обратно в строку
	result := strings.Join(Reverse(splitStr), " ")
	return result
}

func main() {
	fmt.Println(wordReverse("snow dog sun"))
}
