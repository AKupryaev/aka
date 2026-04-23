package main

import (
	"fmt"
	"time"
)

func sender(name string, ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s отправляет: %d\n", name, i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)

	// запускаем две горутины
	go sender("Горутина 1", ch)
	go sender("Горутина 2", ch)

	// читаем 10 значений (по 5 от каждой)
	for i := 0; i < 10; i++ {
		val := <-ch
		fmt.Println("main получил:", val)
	}
}
