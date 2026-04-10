package main

import (
	"fmt"
	"math"
)

func powerFunc(exp int) func(int) int {
	return func(base int) int {
		return int(math.Pow(float64(base), float64(exp)))
	}
}

func main() {
	square := powerFunc(2)
	fmt.Println(square(4)) // 16
	fmt.Println(square(5)) // 25

	cube := powerFunc(3)
	fmt.Println(cube(2)) // 8
	fmt.Println(cube(3)) // 27
}
