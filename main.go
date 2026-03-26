package main

import (
	"errors"
	"fmt"
)

func secondMax(arr []int) (int, error) {
	if len(arr) < 2 {
		return 0, errors.New("меньше двух элементов")
	}

	max, second := arr[0], arr[1]

	if second > max {
		max, second = second, max
	}

	for _, num := range arr[2:] {
		if num > max {
			second = max
			max = num
		} else if num > second && num != max {
			second = num
		}
	}

	return second, nil
}

func main() {
	res, err := secondMax([]int{10, 20, 4, 45, 99})
	if err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println(res)
	}

	res, err = secondMax([]int{5})
	if err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println(res)
	}
}
