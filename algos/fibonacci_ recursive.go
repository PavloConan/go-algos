package algos

import (
	"fmt"
	"time"
)

func FibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FibonacciRecursiveBenchmark(n int) {
	fmt.Println("==== Fibonacci Recursive ====")
	start := time.Now()
	result := FibonacciRecursive(n)
	elapsed := time.Since(start)
	fmt.Printf("Result for %dth item: %d\n", n, result)
	fmt.Printf("Execution time: %d microseconds\n", elapsed.Microseconds())

	fmt.Print("\n")
}
