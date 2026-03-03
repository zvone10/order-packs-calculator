package calculator

// Calculator defines the interface for calculating optimal packaging
type Calculator interface {
	CalculateOptimalPack(numberOfItems int, packageSizes []int) (map[int]int, error)
}
