package main

import (
	"context"
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, t)
	defer cancel()

	if <-ctx.Done(); true {
		return
	}

}

func main() {
	fmt.Println(1)
	sleep(5 * time.Second)
	fmt.Println(2)
}
