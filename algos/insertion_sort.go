package algos

import (
	"fmt"
	"time"
)

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func InsertionSortBenchmark(arr []int) {
	fmt.Println("==== Insertion Sort ====")
	start := time.Now()
	InsertionSort(arr)
	elapsed := time.Since(start)
	fmt.Printf("Execution time for %d items: %d microseconds\n", len(arr), elapsed.Microseconds())

	fmt.Print("\n")
}
