package main

import "fmt"

func safeFunction(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()

	f()
}

func main() {
	safeFunction(func() {
		fmt.Println("Выполняем код...")
		panic("что-то пошло не так")
	})
}
