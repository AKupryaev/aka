package main

import (
	"fmt"
)

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Отправка:", i)
		ch <- i
	}
	close(ch) // закрываем канал после отправки
}

func main() {
	ch := make(chan int)

	go sender(ch)

	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Канал закрыт, выходим")
			break
		}
		fmt.Println("Получено:", val)
	}
}
