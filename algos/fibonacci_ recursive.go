package algos

func FibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}
