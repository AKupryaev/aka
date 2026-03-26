package main

import "fmt"

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	fmt.Println(isLeapYear(2020))
	fmt.Println(isLeapYear(1900))
	fmt.Println(isLeapYear(2000))
}
