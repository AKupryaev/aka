package main

import "fmt"

func sequenceGenerator() func() int {
	counter := 0 // замыкание захватывает переменную counter, которая сохраняет свое состояние между вызовами

	return func() int {
		counter++ // увеличиваем счетчик при каждом вызове функции
		return counter
	}
}

func main() {
	next := sequenceGenerator()

	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3
	fmt.Println(next()) // 4
}
