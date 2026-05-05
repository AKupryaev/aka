package main

import (
	"fmt"
	"time"
)

func main() {

	sem := make(chan struct{}, 3) // три горутины могут работать одновременно

	for i := 0; i < 10; i++ {
		go func(i int) {

			sem <- struct{}{} // заняли слот

			fmt.Println("start", i)
			time.Sleep(1 * time.Second)
			fmt.Println("end", i)

			<-sem // освободили слот
		}(i)
	}

	time.Sleep(5 * time.Second) // ждем, чтобы все горутины завершились
}
