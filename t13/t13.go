package main

import "fmt"

//Меняем местами указатели чисел
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	first := 5
	second := 10
	swap(&first, &second)
	fmt.Println(first, second)
}
