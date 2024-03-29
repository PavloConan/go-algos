package algos

func LinearSearch(m []int, target int) bool {
	for _, n := range m {
		if n == target {
			return true
		}
	}

	return false
}
