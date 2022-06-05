package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]bool)
	for _, value := range s {
		if m[value] {
			return false
		}
		m[value] = true
	}
	return true
}

func main() {
	s := "abCdefAaf"
	b := "abcd"
	fmt.Println(checkUnique(s), checkUnique(b))
}
