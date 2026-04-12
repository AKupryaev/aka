package main

import "fmt"

// Интерфейс сортировщика
type Sorter interface {
	Sort([]int) []int
}

// BubbleSort
type BubbleSort struct{}

func (b BubbleSort) Sort(nums []int) []int {
	// копируем срез, чтобы не менять оригинал
	arr := append([]int(nil), nums...)

	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// QuickSort
type QuickSort struct{}

func (q QuickSort) Sort(nums []int) []int {
	arr := append([]int(nil), nums...)
	return quickSort(arr)
}

// рекурсивная часть quick sort
func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[len(nums)/2]

	var left, middle, right []int

	for _, v := range nums {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			middle = append(middle, v)
		default:
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	result := append(left, middle...)
	result = append(result, right...)

	return result
}

// функция, использующая интерфейс
func sortNumbers(s Sorter, nums []int) []int {
	return s.Sort(nums)
}

func main() {
	bSorter := BubbleSort{}
	qSorter := QuickSort{}

	nums := []int{5, 2, 9, 1, 5, 6}

	fmt.Println(sortNumbers(bSorter, nums)) // [1 2 5 5 6 9]
	fmt.Println(sortNumbers(qSorter, nums)) // [1 2 5 5 6 9]
}
