package calculator

import "errors"

// Error variables for the calculator package
var (
	ErrDuplicatePackageSizes = errors.New("package sizes contain duplicate values")
)

// Calculator defines the interface for calculating optimal packaging
type Calculator interface {
	CalculateOptimalPack(numberOfItems int, packageSizes []int) (map[int]int, error)
}
