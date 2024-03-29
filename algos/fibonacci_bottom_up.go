package algos

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
