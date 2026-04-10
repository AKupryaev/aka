package main

import "fmt"

func removeAtIndex(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	fmt.Println(removeAtIndex([]int{1, 2, 3, 4, 5}, 2)) // [1 2 4 5]
	fmt.Println(removeAtIndex([]int{10, 20, 30}, 5))    // [10 20 30]
}
