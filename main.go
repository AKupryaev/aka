package main

import (
	"fmt"
	"sync"
)

// --- Stage 1: генератор (pipeline) ---
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// --- Worker (fan-out) ---
func worker(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("Worker %d обрабатывает %d\n", id, n)
			out <- n * n // например, квадрат
		}
		close(out)
	}()
	return out
}

// --- Fan-In ---
func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// --- Stage 2: финальная обработка (pipeline) ---
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func main() {
	// pipeline start
	input := gen(1, 2, 3, 4, 5)

	// fan-out (3 воркера)
	w1 := worker(1, input)
	w2 := worker(2, input)
	w3 := worker(3, input)

	// fan-in
	merged := fanIn(w1, w2, w3)

	// pipeline end
	result := double(merged)

	// вывод
	for v := range result {
		fmt.Println("Результат:", v)
	}
}
