package main

import (
	"fmt"
	"sync"
)

//Структура счетчик с использованием мьютекса для того, чтобы получить корректный результат в конкурентной среде
type Counter struct {
	number int
	mu     sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.number++
}

func (c *Counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.number--
}

func main() {
	c := &Counter{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 200; i++ {
		wg.Add(1)

		go func(num int, cnt *Counter) {
			defer wg.Done()
			cnt.Increment()
		}(i, c)
	}

	wg.Wait()
	fmt.Println(c.number)
}
