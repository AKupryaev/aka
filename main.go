package main

import "fmt"

func invertMap(m map[string]int) map[int]string {
	inverted := make(map[int]string)

	for key, value := range m {
		inverted[value] = key
	}

	return inverted
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(invertMap(m)) // map[1:"a" или "c", 2:"b"]
}
