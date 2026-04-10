package main

import (
	"fmt"
	"strconv"
)

func reverseNumber(n int) int {
	isNegative := n < 0
	if isNegative {
		n = -n
	}

	str := strconv.Itoa(n)

	bytes := []byte(str)

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	reversedStr := string(bytes)

	result, _ := strconv.Atoi(reversedStr)

	if isNegative {
		result = -result
	}

	return result
}

func main() {
	fmt.Println(reverseNumber(123))  // 321
	fmt.Println(reverseNumber(-123)) // -123
}
