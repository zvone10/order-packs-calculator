package calculator

import (
	"sort"
)

// Calculator defines the interface for calculating optimal packaging
type Calculator interface {
	CalculateOptimalPack(numberOfItems int, packageSizes []int) (map[int]int, error)
}

type calculator struct{}

// NewCalculator creates a new instance of the calculator
func NewCalculator() Calculator {
	return &calculator{}
}

func (c *calculator) CalculateOptimalPack(numberOfItems int, packageSizes []int) (map[int]int, error) {
	remainingItems := numberOfItems

	// sort package sizes in descending order to maximize the use of larger packages
	sort.Sort(sort.Reverse(sort.IntSlice(packageSizes)))

	result := make(map[int]int)
	for i, capacity := range packageSizes {

		boxes := remainingItems / capacity
		if boxes > 0 {
			result[capacity] = boxes
			remainingItems -= boxes * capacity
		}

		if boxes == 0 && i == len(packageSizes)-1 && remainingItems > 0 {
			// if we are at the smallest package and still have remaining items, we need one more box
			result[capacity] = 1
			remainingItems = 0
		}
	}
	return result, nil
}
