package main

import (
	"fmt"
	"time"
)

func sender(name string, ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s отправляет: %d\n", name, i)
		ch <- i // не блокируется, пока буфер не заполнен
	}
	fmt.Println(name, "завершил отправку")
}

func main() {
	// буфер на 3 элемента
	ch := make(chan int, 3)

	go sender("Горутина 1", ch)
	go sender("Горутина 2", ch)

	time.Sleep(1 * time.Second) // даём горутинам заполнить буфер

	for i := 0; i < 10; i++ {
		val := <-ch
		fmt.Println("main получил:", val)
		time.Sleep(300 * time.Millisecond)
	}
}
