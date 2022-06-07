package main

import "fmt"

// Определяем структуру Human
type Human struct {
	Name   string
	Gender string
	Age    int
}

//Реализуем встраивание структуры Human в структуру Action
type Action struct {
	Human
}

//Конструктор Human
func (h *Human) newHuman(name, gender string, age int) {
	h.Name = name
	h.Gender = gender
	h.Age = age
}

func main() {
	//Определили Action и через него можем использовать методы и поля Human
	human := Action{}
	human.newHuman("Alexander", "man", 22)
	fmt.Println(human)
}
