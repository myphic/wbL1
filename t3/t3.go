package main

import (
	"fmt"
	"sync"
)

func main() {
	//Создаем массив с самовычисляемой длиной
	arr := [...]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	sum := 0
	//Чтобы дождаться выполнения горутин используем WaitGroup
	for _, value := range arr {
		wg.Add(1)
		go func(value int) {
			sum += value * value
			defer wg.Done()
		}(value)
	}
	wg.Wait()
	fmt.Println(sum)
}
