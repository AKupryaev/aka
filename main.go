package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := range 10000 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for i := range 10000 {
			ch <- i * 2
		}
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Println("v =", v, "worker1")
		}
	}()

	for v := range ch {
		fmt.Println("v =", v, "worker2")
	}

}
