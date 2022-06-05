package main

import "fmt"

func strReverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, r := range s {
		n--
		runes[n] = r
	}

	return string(runes[n:])
}

func main() {
	fmt.Println(strReverse("главрыба"))
}
