package main

import (
	"fmt"
)

func recursive_bubble_sort(arr []int, size int) []int {
	if size == 1 {
		return arr
	}

	var i = 0
	for i < size-1 {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}

		i++
	}

	recursive_bubble_sort(arr, size-1)

	return arr
}

func main() {
	data := []int{3, 7, 2, 6, 6, 1, 5}
	r := recursive_bubble_sort(data, len(data))
	fmt.Println(r)
}
