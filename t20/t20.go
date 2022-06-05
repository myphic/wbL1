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
	splitStr := strings.Split(s, " ")
	result := strings.Join(Reverse(splitStr), " ")
	return result
}

func main() {
	fmt.Println(wordReverse("snow dog sun"))
}
