package algos

func BinarySearch(m []int, target int) bool {
	hi := len(m) - 1
	lo := 0

	return search(m, hi, lo, target)
}

func search(m []int, hi int, lo int, target int) bool {
	if lo > hi {
		return false
	}

	idx := (lo + hi) / 2

	if m[idx] == target {
		return true
	} else if m[idx] > target {
		return search(m,m[idx-1], lo, target)
	} else {
		return search(m, hi, m[idx+1], target)
	}
}
