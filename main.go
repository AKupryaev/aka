package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// случайное время от 0 до 2 секунд
	duration := time.Duration(rand.Intn(2000)) * time.Millisecond

	fmt.Printf("Worker %d старт, будет работать %v\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Worker %d завершился\n", id)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup
	N := 5

	wg.Add(N)

	for i := 1; i <= N; i++ {
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Все горутины завершены")
}
