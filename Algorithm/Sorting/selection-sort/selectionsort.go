package main

import "fmt"

func main() {
	arr := []int{1, 39, 2, 9, 7, 54, 11}
	n := len(arr)

	for i := 1; i < n-1; i++ {
		j := i + 1
		minIndex := i

		if j < n {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			j++
		}

		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
		i++
	}

	fmt.Println(arr)
}
