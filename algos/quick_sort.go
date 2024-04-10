package algos

import (
	"fmt"
	"time"
)

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)
	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func QuickSortBenchmark(arr []int) {
	fmt.Println("==== Quick Sort ====")
	start := time.Now()
	QuickSort(arr)
	elapsed := time.Since(start)
	fmt.Printf("Execution time for %d items: %d microseconds\n", len(arr), elapsed.Microseconds())

	fmt.Print("\n")
}
