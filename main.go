package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
		ch2 <- 2
	}()

	for i := 0; i < 2; i++ {

		select {
		case v := <-ch1:
			fmt.Println("v =", v, "from ch1")
		case v := <-ch2:
			fmt.Println("v =", v, "from ch2")

		}
	}
}
