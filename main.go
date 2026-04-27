package main

import (
	"fmt"
	"time"
)

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i + 1
		}
		close(ch)
	}()

	return ch
}

func double(inputCh <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for v := range inputCh {
			time.Sleep(500 * time.Millisecond)
			ch <- v * 2
		}
		close(ch)
	}()

	return ch
}

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	reader(double(writer()))
}
