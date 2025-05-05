package main

import (
	"fmt"
	"time"
)

const max int = 50

func main() {
	t := time.Now().UnixMicro()

	jobs := make(chan int, max)
	results := make(chan int, max)

	go worker(jobs, results)

	for i := 0; i < max; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < max; i++ {
		fmt.Printf("result: %d\n", <-results)
	}
	defer close(results)

	fmt.Printf("Took %v ms\n", float64(time.Now().UnixMicro() - t) / 1000.00)
}

func worker(jobs <-chan int, result chan <- int){
	for j := range jobs {
		result <- fib(j)
	}
}

func fib(num int) int {
	// Add memoisation for optimization
	cache := map[int]int{
		0 : 1,
		1 : 1,
	}

	var helper func(n int)int

	helper = func(n int)int {
		if n <= 1 {
			return n
		}

		// Check if we have the answer before calling recursion
		if v, ok := cache[n]; ok {
			return v
		}

		// Call recursion stack
		ans := helper(n - 1) + helper(n - 2)
		
		// Cache response
		cache[n] = ans
		
		// Return result
		return ans
	}

	return helper(num)
}

