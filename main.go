package main

import "fmt"

func calculator(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return 0
		}
		return a / b
	default:
		return 0
	}
}

func main() {
	fmt.Println(calculator(10, 5, "+")) // 15
	fmt.Println(calculator(10, 5, "-")) // 5
	fmt.Println(calculator(10, 5, "*")) // 50
	fmt.Println(calculator(10, 0, "/")) // 0
}
