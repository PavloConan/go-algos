package main

import (
	"fmt"
	"go-algos/algos"
	"math/rand"
	"time"
)

type intSlice []int

func (s intSlice) Shuffle() {
	for i := range s {
		r := rand.Intn(i + 1)
		s[i], s[r] = s[r], s[i]
	}
}

func main() {
	fmt.Println("Fibonacci Recursive")
	start := time.Now()
	result := algos.FibonacciRecursive(40)
	elapsed := time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Execution time: %d microseconds\n", elapsed.Microseconds())

	fmt.Println("Fibonacci Memo")
	start = time.Now()
	result = algos.FibonacciMemo(40)
	elapsed = time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Execution time: %d microseconds\n", elapsed.Microseconds())

	fmt.Println("Fibonacci Bottom Up")
	start = time.Now()
	result = algos.FibonacciBottomUp(40)
	elapsed = time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Execution time: %d microseconds\n", elapsed.Microseconds())

	sortedArr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		sortedArr[i] = i
	}

	fmt.Println("Binary Search")
	start = time.Now()
	found := algos.BinarySearch(sortedArr, 900000)
	elapsed = time.Since(start)
	fmt.Printf("Result: %t\n", found)
	fmt.Printf("Execution time: %d microseconds\n", elapsed.Microseconds())

	fmt.Println("Linear Search")
	start = time.Now()
	found = algos.LinearSearch(sortedArr, 900000)
	elapsed = time.Since(start)
	fmt.Printf("Result: %t\n", found)
	fmt.Printf("Execution time: %d microseconds\n", elapsed.Microseconds())

	unsortedArr := make(intSlice, 1000000)
	for i := 0; i < 1000000; i++ {
		unsortedArr[i] = 1000000 - i
	}

	unsortedArr.Shuffle()
}
