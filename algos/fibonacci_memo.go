package algos

import (
	"fmt"
	"time"
)

func FibonacciMemo(n int) int {
	m := make(map[int]int)
	return memoRecursive(n, m)
}

func memoRecursive(n int, m map[int]int) int {
	if n < 2 {
		return n
	}

	if _, ok := m[n]; !ok {
		m[n] = memoRecursive(n-1, m) + memoRecursive(n-2, m)
	}

	return m[n]
}

func FibonacciMemoBenchmark(n int) {
	fmt.Println("==== Fibonacci Memo ====")
	start := time.Now()
	result := FibonacciMemo(n)
	elapsed := time.Since(start)
	fmt.Printf("Result for %dth item: %d\n", n, result)
	fmt.Printf("Execution time: %d microseconds\n", elapsed.Microseconds())

	fmt.Print("\n")
}
