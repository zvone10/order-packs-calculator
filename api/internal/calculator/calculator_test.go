package calculator

import (
	"testing"
)

func TestCalculateOptimalPack(t *testing.T) {
	calc := NewKnapsackCalculator()

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
		{
			name:          "12001 items with capacities [250, 500, 1000, 2000, 5000]",
			numberOfItems: 12001,
			boxCapacity:   []int{250, 500, 1000, 2000, 5000},
			expectedResult: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
			expectError: false,
		},
		{
			name:          "500000 items with capacities [23, 31, 53]",
			numberOfItems: 500000,
			boxCapacity:   []int{23, 31, 53},
			expectedResult: map[int]int{
				53: 9429,
				31: 7,
				23: 2,
			},
			expectError: false,
		},
		{
			name:          "15 items with capacities [5, 7]",
			numberOfItems: 15,
			boxCapacity:   []int{5, 7},
			expectedResult: map[int]int{
				5: 3,
			},
			expectError: false,
		},
		{
			name:          "Large target",
			numberOfItems: 5000000,
			boxCapacity:   []int{17, 65, 88, 111},
			expectedResult: map[int]int{
				111: 45044,
				17:  3,
				65:  1,
			},
			expectError: false,
		},
		{
			name:          "Large target which is multiple",
			numberOfItems: 520000,
			boxCapacity:   []int{5, 26, 53},
			expectedResult: map[int]int{
				5:  4,
				26: 6,
				53: 9808,
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
