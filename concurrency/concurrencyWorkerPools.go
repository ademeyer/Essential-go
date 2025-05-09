package main

import (
	"fmt"
	"time"
)

const max int = 15

func main() {
	t := time.Now().UnixMicro()

	jobs := make(chan int, max)
	results := make(chan float64, max)

	go worker(jobs, results)
  go worker(jobs, results)
  go worker(jobs, results)

	for i := 0; i < max; i++ {
		jobs <- i
	}
	close(jobs)
  
  twoSeries := [2]float64{0, 0}
	for i := 0; i < max; i++ {
    twoSeries[i % 2] = <-results
	}
	defer close(results)

  fmt.Printf("total result: %v, and took %v ms\n", (twoSeries[0] + twoSeries[1]), float64(time.Now().UnixMicro() - t) / 1000.00)
}

func worker(jobs <-chan int, result chan <- float64){
	for j := range jobs {
		result <- float64(fib(j))
	}
}

func fib(num int) int {
	// Add memoisation for optimization
	cache := map[int]int{
    0 : 0,
    1 : 1,
    2 : 1,
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

