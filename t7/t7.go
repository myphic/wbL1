package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i + 1
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}
