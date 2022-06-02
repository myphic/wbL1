package main

import (
	"fmt"
	"sync"
)

func main() {
	//WaitGroup для того чтобы дождаться конца выполнения горутин
	var wg sync.WaitGroup
	wg.Add(3)

	c1 := make(chan int)
	c2 := make(chan int)
	numbers := []int{2, 4, 6, 8}
	go sender(c1, numbers, &wg)
	go doubleNumber(c1, c2, &wg)
	go printer(c2, &wg)

	wg.Wait()
}

//Конвейер горутин, sender отправляет, doubleNumber удваивает, printer печатает на экран
func sender(outputStream chan int, numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range numbers {
		outputStream <- v
	}
	close(outputStream)
}

func doubleNumber(inputStream, outputStream chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range inputStream {
		outputStream <- v * 2
	}
	close(outputStream)
}

func printer(inputStream chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range inputStream {
		fmt.Println(v)
	}
}
