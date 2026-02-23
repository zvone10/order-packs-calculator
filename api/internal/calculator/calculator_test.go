package calculator

import (
	"testing"
)

func TestCalculateOptimalPack(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name           string
		numberOfItems  int
		boxCapacity    []int
		expectedResult map[int]int
		expectError    bool
	}{
		{
			name:          "501 items with capacities [250, 500, 1000, 2000, 5000]",
			numberOfItems: 501,
			boxCapacity:   []int{250, 500, 1000, 2000, 5000},
			expectedResult: map[int]int{
				250: 1,
				500: 1,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.CalculateOptimalPack(tt.numberOfItems, tt.boxCapacity)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err != nil)
			}

			if len(result) != len(tt.expectedResult) {
				t.Errorf("expected result length: %d, got: %d", len(tt.expectedResult), len(result))
			}

			for capacity, count := range tt.expectedResult {
				if result[capacity] != count {
					t.Errorf("for capacity %d, expected count %d, got %d", capacity, count, result[capacity])
				}
			}
		})
	}
}
