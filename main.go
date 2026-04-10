package main

import "fmt"

func mergeMaps(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)

	// добавляем элементы из первой map
	for key, value := range m1 {
		result[key] = value
	}

	// добавляем элементы из второй map (суммируем при совпадении)
	for key, value := range m2 {
		result[key] += value
	}

	return result
}

func main() {
	m1 := map[string]int{"apple": 5, "banana": 3}
	m2 := map[string]int{"banana": 2, "orange": 4}

	fmt.Println(mergeMaps(m1, m2)) // map[apple:5 banana:5 orange:4]
}
