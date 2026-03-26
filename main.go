package main

import "fmt"

func maxOfThree(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}

func main() {

	fmt.Println(maxOfThree(3, 7, 5))  // Output: 7
	fmt.Println(maxOfThree(10, 2, 8)) // Output: 10

}
