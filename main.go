package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordCount(text string) map[string]int {
	result := make(map[string]int)

	// Приводим к нижнему регистру
	text = strings.ToLower(text)

	// Удаляем знаки препинания
	clean := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, text)

	// Разбиваем на слова
	words := strings.Fields(clean)

	// Считаем
	for _, word := range words {
		result[word]++
	}

	return result
}

func main() {
	fmt.Println(wordCount("Hello, hello world! Hello world."))
	// map[hello:3 world:2]
}
