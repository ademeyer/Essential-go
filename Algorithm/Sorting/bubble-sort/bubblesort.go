package main

import (
	"fmt"
)

func bubble_sort(data []int) []int {
	n := len(data)
	isDone := false
	for !isDone {
		isDone = true
		var i = 0
		for i < n-1 {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				isDone = false
			}
			i++
		}
	}
	return data
}

func main() {
	data := []int{3, 7, 1, 2, 6, 6, 1, 2, 5}
	r := bubble_sort(data)
	fmt.Println(r)
}
