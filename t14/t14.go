package main

import (
	"fmt"
	"reflect"
)

//Используем reflect потому что switch type не может определить тип для канала
func getType(variable interface{}) string {
	v := reflect.ValueOf(variable)
	switch v.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Chan:
		return "chan"
	default:
		fmt.Println("Unknown type")
	}
	return ""
}

func main() {
	c := make(chan int)
	arr := []int{1}
	b := true
	fmt.Println(getType(arr))
	fmt.Println(getType(c))
	fmt.Println(getType(b))
}
