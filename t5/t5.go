package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	timeValue := 5
	//Создаем контекст с таймаутом чтобы завершиться по истечению timeValue секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeValue))
	defer cancel()

	messages := make(chan int)
	//Пишем сообщения в канал messages
	go func() {
		for i := 0; ; i++ {
			messages <- i
		}
	}()
L:
	for {
		select {
		//Читаем с канала и выводим в stdout
		case msg := <-messages:
			fmt.Println(msg)
		//По истечению времени выходим из цикла
		case <-ctx.Done():
			break L
		}
	}
}
