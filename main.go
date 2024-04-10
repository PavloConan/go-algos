package main

import (
	"math/rand"
)

type intSlice []int

func (s intSlice) Shuffle() {
	for i := range s {
		r := rand.Intn(i + 1)
		s[i], s[r] = s[r], s[i]
	}
}

func main() {
	// 	algos.FibonacciRecursiveBenchmark(40)
	// 	algos.FibonacciMemoBenchmark(40)
	// 	algos.FibonacciBottomUpBenchmark(40)

	// 	sortedArr := make([]int, 1000000)
	// 	for i := 0; i < 1000000; i++ {
	// 		sortedArr[i] = i
	// 	}

	// 	algos.LinearSearchBenchmark(sortedArr, 900000)
	// 	algos.BinarySearchBenchmark(sortedArr, 900000)

	// 	unsortedArr := make(intSlice, 5000)
	// 	for i := 0; i < 5000; i++ {
	// 		unsortedArr[i] = i
	// 	}

	// 	unsortedArr.Shuffle()

	// 	algos.BubbleSortBenchmark(unsortedArr)

	// 	unsortedArr.Shuffle()

	// 	algos.QuickSortBenchmark(unsortedArr)

	// 	unsortedArr.Shuffle()

	// 	algos.MergeSortBenchmark(unsortedArr)

	// 	unsortedArr.Shuffle()

	// 	algos.InsertionSortBenchmark(unsortedArr)
}
