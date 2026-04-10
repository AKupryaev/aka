package main

import "fmt"

func increment(prt *int) {
	*prt++
}

func main() {
	n := 5
	increment(&n)
	fmt.Println(n) // 6
}
