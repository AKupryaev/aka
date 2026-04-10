package main

import "fmt"

func safeDivide(a, b int) (result int) {
	// defer с recover
	defer func() {
		if r := recover(); r != nil {
			result = 0
		}
	}()

	if b == 0 {
		panic("division by zero")
	}

	return a / b
}

func main() {
	fmt.Println(safeDivide(10, 2)) // 5
	fmt.Println(safeDivide(10, 0)) // 0
}
