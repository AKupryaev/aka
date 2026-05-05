package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done() // забыли Done() без дефера зависнет навсегда
		fmt.Println("работаю")
	}()

	wg.Wait() // ✅ не зависнет навсегда
}
