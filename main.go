package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup

	const goroutines = 10
	const ops = 1000

	// запись (writers)
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < ops; j++ {
				m.Store(fmt.Sprintf("goroutine-%d-key-%d", id, j), j)
			}
		}(i)
	}

	// чтение (readers)
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < ops; j++ {
				key := fmt.Sprintf("goroutine-%d-key-%d", id, j)
				value, ok := m.Load(key)
				if ok {
					_ = value
				}
			}
		}(i)
	}

	wg.Wait()

	// итоговая проверка
	count := 0
	m.Range(func(key, value any) bool {
		count++
		return true
	})

	fmt.Println("Всего элементов:", count)
}
