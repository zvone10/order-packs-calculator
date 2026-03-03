package calculator

import (
	"sort"
)

type KnapsackCalculator struct{}

// NewKnapsackCalculator creates a new instance of the KnapsackCalculator
func NewKnapsackCalculator() Calculator {
	return &KnapsackCalculator{}
}

func (c *KnapsackCalculator) CalculateOptimalPack(numberOfItems int, packageSizes []int) (map[int]int, error) {
	//if there are equal items in packageSizes return error
	if HasSameItems(packageSizes) {
		return nil, ErrDuplicatePackageSizes
	}
	sort.Ints(packageSizes)
	minSize, maxSize := packageSizes[0], packageSizes[len(packageSizes)-1]

	memoization := make(map[int]int)
	parent := make(map[int]int)

	memoization[0] = 0
	for _, s := range packageSizes {
		memoization[s] = 1
		parent[s] = s
	}

	//knapsack algorithm
	for i := minSize; i <= numberOfItems+maxSize; i++ {
		for _, s := range packageSizes {
			if v, ok := memoization[i-s]; ok && v != 0 && i >= s {
				initialValue, initialExists := memoization[i]
				if initialExists && v+1 < initialValue {
					memoization[i] = v + 1
					parent[i] = s
				} else if !initialExists {
					memoization[i] = v + 1
					parent[i] = s
				}
			}
		}
	}

	minimalNumberOfPackages := -1
	for i := numberOfItems; i <= numberOfItems+maxSize; i++ {
		if v, ok := memoization[i]; ok && v != 0 {
			minimalNumberOfPackages = i
			break
		}
	}

	optimal := make(map[int]int)
	for minimalNumberOfPackages > 0 {
		p := parent[minimalNumberOfPackages]
		optimal[p]++
		minimalNumberOfPackages -= p
	}

	return optimal, nil
}
