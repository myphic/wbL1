package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(countWorkers int, sign chan os.Signal) {
	wg := sync.WaitGroup{}
	container := make(chan int, countWorkers)

	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 0; i < countWorkers; i++ {
		container <- i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case value := <-container:
					fmt.Println(value)
					container <- value
				case <-ctx.Done():
					return
				}
			}
		}()

	}
	<-sign
	fmt.Println("--- SHUTDOWN SIGNAL RECEIVED ---")
	cancelFunc()
	wg.Wait()
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	worker(4, sigs)
	fmt.Println("exiting")
}
