package main

import "fmt"

func charCase(r rune) string {
	switch {
	case r >= 'A' && r <= 'Z':
		return "Латинская заглавная"
	case r >= 'a' && r <= 'z':
		return "Латинская строчная"
	case r >= 'А' && r <= 'Я':
		return "Кириллическая заглавная"
	case r >= 'а' && r <= 'я':
		return "Кириллическая строчная"
	default:
		return "Другое"
	}
}

func main() {
	fmt.Println(charCase('A')) // Латинская заглавная
	fmt.Println(charCase('я')) // Кириллическая строчная
	fmt.Println(charCase('Б')) // Кириллическая заглавная
	fmt.Println(charCase('1')) // Другое
	fmt.Println(charCase('😊')) // Другое
}
