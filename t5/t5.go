package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeValue := 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeValue))
	defer cancel()

	messages := make(chan int)

	go func() {
		for i := 0; ; i++ {
			messages <- i
		}
	}()
L:
	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		case <-ctx.Done():
			break L
		}
	}
}
