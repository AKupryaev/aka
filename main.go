package main

import "fmt"

func sumTon(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
func main() {

	fmt.Println(sumTon(5))  //15
	fmt.Println(sumTon(10)) //55
}
