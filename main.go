package main

import (
	"fmt"
	"sort"
)

func mergeAndSort(arr1, arr2 []int) []int {

	result := append(arr1, arr2...)

	sort.Ints(result)

	return result
}

func main1() {
	fmt.Println(mergeAndSort([]int{3, 1, 5}, []int{4, 2, 6})) // [1 2 3 4 5 6]
	fmt.Println(mergeAndSort([]int{8, 2, 0}, []int{7, 3, 1})) // [0 1 2 3 7 8]
}
