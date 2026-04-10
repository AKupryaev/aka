package main

import "fmt"

func appendValue(slice *[]int, value int) {
	*slice = append(*slice, value)
}

func main() {
	nums := []int{1, 2, 3}
	appendValue(&nums, 4)
	fmt.Println(nums) // [1 2 3 4]
}
