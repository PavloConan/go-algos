package structures

import (
	"testing"
)

// function to test queue Peek function
func TestQueue_Peek(t *testing.T) {
	tests := []struct {
		name     string
		queue    Queue
		expected int
	}{
		{
			name: "empty queue",
			queue: Queue{
				Length: 0,
				Head:   nil,
				Tail:   nil,
			},
			expected: 0,
		},
		{
			name: "non-empty queue",
			queue: Queue{
				Length: 3,
				Head: &QNode{
					Value: 1,
					Next: &QNode{
						Value: 2,
						Next: &QNode{
							Value: 3,
							Next:  nil,
						},
					},
				},
				Tail: &QNode{
					Value: 3,
					Next:  nil,
				},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.queue.Peek()
			if actual != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
