package algos

import (
	"fmt"
	"time"
)

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSortBenchmark(arr []int) {
	fmt.Println("==== Bubble Sort ====")
	start := time.Now()
	BubbleSort(arr)
	elapsed := time.Since(start)
	fmt.Printf("Execution time for %d items: %d microseconds\n", len(arr), elapsed.Microseconds())

	fmt.Print("\n")
}
