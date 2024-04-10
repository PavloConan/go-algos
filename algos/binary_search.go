package algos

import (
	"fmt"
	"time"
)

func BinarySearch(arr []int, target int) bool {
	hi := len(arr) - 1
	lo := 0

	return search(arr, hi, lo, target)
}

func search(arr []int, hi int, lo int, target int) bool {
	if lo > hi {
		return false
	}

	idx := (lo + hi) / 2

	if arr[idx] == target {
		return true
	} else if arr[idx] > target {
		return search(arr, idx-1, lo, target)
	} else {
		return search(arr, hi, idx+1, target)
	}
}

func BinarySearchBenchmark(arr []int, target int) {
	fmt.Println("==== Binary Search ====")
	start := time.Now()
	found := BinarySearch(arr, target)
	elapsed := time.Since(start)
	fmt.Printf("Result: %t\n", found)
	fmt.Printf("Execution time for %d items: %d microseconds\n", len(arr), elapsed.Microseconds())

	fmt.Print("\n")
}
