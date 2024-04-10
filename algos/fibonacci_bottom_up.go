package algos

import (
	"fmt"
	"time"
)

func FibonacciBottomUp(n int) int {
	if n < 2 {
		return n
	}

	fib := []int{0, 1}

	for i := 2; i <= n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib[n]
}

func FibonacciBottomUpBenchmark(n int) {
	fmt.Println("==== Fibonacci Bottom Up ====")
	start := time.Now()
	result := FibonacciBottomUp(n)
	elapsed := time.Since(start)
	fmt.Printf("Result for %dth item: %d\n", n, result)
	fmt.Printf("Execution time: %d microseconds\n", elapsed.Microseconds())

	fmt.Print("\n")
}
