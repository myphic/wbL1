package main

import "fmt"

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5.0}
	result := map[int][]float32{}
	for _, temp := range temps {
		//определяем ключ и пушим в мапу
		key := (int(temp) / 10) * 10
		result[key] = append(result[key], temp)
	}

	for k, v := range result {
		fmt.Println(k, v)
	}
}
