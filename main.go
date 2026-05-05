package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Value string
}

var (
	once     sync.Once
	instance *Config
)

func getConfig() *Config {
	once.Do(func() {
		fmt.Println("Инициализация...")
		instance = &Config{
			Value: "Настройки загружены",
		}
	})
	return instance
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			cfg := getConfig()
			fmt.Printf("Goroutine %d: %s\n", id, cfg.Value)
		}(i)
	}

	wg.Wait()
}
