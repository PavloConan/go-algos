package algos

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "sorted",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "reversed",
			arr:      []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "random",
			arr:      []int{5, 2, 3, 4, 1, 6, 7, 8, 10, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.arr)
			if !compareSlices(tt.arr, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, tt.arr)
			}
		})
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
