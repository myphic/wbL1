package main

import "fmt"

func main() {
	first := 5
	second := 10
	first = first + second
	second = first - second
	first = first - second
	fmt.Println(first, second)
}
