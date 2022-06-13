package main

import "fmt"

/*
	Адаптер - паттерн проектирования, который позволяет объектам с несовместимыми интерфейсами работать вместе.
*/
// Неизменный внешний сервис
type WeatherService interface {
	Request() int32
}
type Weather struct {
	value int32
}

func (w *Weather) Request() int32 {
	return w.value
}

// /Неизменный внешний сервис

type MyService struct {
	weatherValue float32
}

type MyServiceAdapter struct {
	myService *MyService
}

func (service *MyServiceAdapter) Request() int32 {
	return int32(service.myService.weatherValue)
}

func main() {
	m := &MyServiceAdapter{&MyService{32.4}}
	req := m.Request()
	fmt.Print(req)
}
