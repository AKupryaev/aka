package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Структура
type TextProcessor struct {
	Text string
}

// Подсчёт слов
func (tp TextProcessor) WordCount() map[string]int {
	result := make(map[string]int)

	// приводим к нижнему регистру
	text := strings.ToLower(tp.Text)

	// удаляем знаки препинания
	clean := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, text)

	// разбиваем на слова
	words := strings.Fields(clean)

	// считаем
	for _, w := range words {
		result[w]++
	}

	return result
}

// Замена слов
func (tp *TextProcessor) ReplaceWord(old, new string) {
	// делаем замену без учёта регистра (упрощённо)
	words := strings.Fields(tp.Text)

	for i, w := range words {
		// убираем знаки препинания для сравнения
		clean := strings.Trim(w, ".,!?")

		if strings.EqualFold(clean, old) {
			// сохраняем пунктуацию
			prefix := strings.TrimSuffix(w, clean)
			suffix := strings.TrimPrefix(w, clean)

			words[i] = prefix + new + suffix
		}
	}

	tp.Text = strings.Join(words, " ")
}

func main() {
	tp := TextProcessor{"Hello world, hello again!"}

	fmt.Println(tp.WordCount()) // map[hello:2 world:1 again:1]

	tp.ReplaceWord("hello", "hi")
	fmt.Println(tp.Text) // Hi world, hi again!
}
