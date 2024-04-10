package algos

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected bool
	}{
		{
			name:     "found",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   5,
			expected: true,
		},
		{
			name:     "not found",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   11,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := LinearSearch(tt.arr, tt.target)
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
