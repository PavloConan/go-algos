package algos

import (
	"fmt"
	"time"
)

func MergeSort(arr []int) {
	ms(arr, 0, len(arr)-1)
}

func ms(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2

	ms(arr, lo, mid)
	ms(arr, mid+1, hi)

	merge(arr, lo, mid, hi)
}

func merge(arr []int, lo, mid, hi int) {
	left := make([]int, mid-lo+1)
	right := make([]int, hi-mid)

	for i := 0; i < len(left); i++ {
		left[i] = arr[lo+i]
	}

	for i := 0; i < len(right); i++ {
		right[i] = arr[mid+1+i]
	}

	i, j, k := 0, 0, lo

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < mid-lo+1 {
		arr[k] = left[i]
		i++
		k++
	}

	for j < hi-mid {
		arr[k] = right[j]
		j++
		k++
	}
}

func MergeSortBenchmark(arr []int) {
	fmt.Println("==== Merge Sort ====")
	start := time.Now()
	MergeSort(arr)
	elapsed := time.Since(start)
	fmt.Printf("Execution time for %d items: %d microseconds\n", len(arr), elapsed.Microseconds())

	fmt.Print("\n")
}
