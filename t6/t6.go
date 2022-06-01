package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func firstWay() {
	// 1 способ через bool канал, когда передаем true в quit, то осуществляется выход из горутины
	quit := make(chan bool)
	go func() {
		fmt.Println("Text in goroutine")
		for {
			select {
			case <-quit:
				return
			}
		}
	}()
	fmt.Println("Text after goroutine")
	quit <- true
}

func secondWay() {
	// 2 способ уничтожить main горутину
	go func() {
		fmt.Println("Text in goroutine")
		os.Exit(1)
	}()
	fmt.Println("Text after goroutine")
	time.Sleep(time.Second * 5)
}

func thirdWay() {
	//Через канал и WaitGroup
	var wg sync.WaitGroup
	c := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Text in goroutine")
		<-c
	}()
	c <- true
	fmt.Println("Text after goroutine")
	close(c)
	wg.Wait()
}

func fourthWay() {
	//С помощью for .. range для каналов, можно поменять на v, ok := <-c и делать проверку на !ok
	ch := make(chan bool)
	go func() {
		fmt.Println("Text in goroutine")
		for {
			for v := range ch {
				fmt.Println(v)
			}
		}
	}()
	fmt.Println("Text after goroutine")
	ch <- true
	close(ch)
	time.Sleep(time.Second)
}

func fifthWay() {
	//Используем context withcancel
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		fmt.Println("Text in goroutine")
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			}
		}
	}(ctx)
	fmt.Println("Text after goroutine")
	cancel()
	<-ch

}

func sixthWay() {
	//Используя context WithTimeout
	ch := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	go func(ctx context.Context) {
		fmt.Println("Text in goroutine")
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			}
		}
	}(ctx)
	fmt.Println("Text after goroutine")
	defer cancel()
	<-ch
}

func main() {
	sixthWay()
}
