package algos

import (
	"fmt"
	"time"
)

func LinearSearch(arr []int, target int) bool {
	for _, n := range arr {
		if n == target {
			return true
		}
	}

	return false
}

func LinearSearchBenchmark(arr []int, target int) {
	fmt.Println("==== Linear Search ====")
	start := time.Now()
	found := LinearSearch(arr, target)
	elapsed := time.Since(start)
	fmt.Printf("Result: %t\n", found)
	fmt.Printf("Execution time for %d items: %d microseconds\n", len(arr), elapsed.Microseconds())

	fmt.Print("\n")
}
