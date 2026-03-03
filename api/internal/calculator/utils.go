package calculator

import "testing"

func HasSameItems(items []int) bool {
	seen := make(map[int]bool)
	for _, item := range items {
		if seen[item] {
			return true
		}
		seen[item] = true
	}
	return false
}

//unit tests for HasSameItems function

func TestHasSameItems(t *testing.T) {
	tests := []struct {
		name     string
		items    []int
		expected bool
	}{
		{
			name:     "No duplicates",
			items:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "With duplicates",
			items:    []int{1, 2, 3, 4, 5, 2},
			expected: true,
		},
		{
			name:     "Empty slice",
			items:    []int{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasSameItems(tt.items)
			if result != tt.expected {
				t.Errorf("Expected %v but got %v", tt.expected, result)
			}
		})
	}
}
