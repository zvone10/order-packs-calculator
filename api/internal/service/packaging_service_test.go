package service

import (
	"errors"
	"testing"

	"order-pack-calculator-api/internal/calculator"
	"order-pack-calculator-api/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCalculator is a mock implementation of the Calculator interface
type MockCalculator struct {
	mock.Mock
}

func (m *MockCalculator) CalculateOptimalPack(numberOfItems int, packageSizes []int) (map[int]int, error) {
	args := m.Called(numberOfItems, packageSizes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[int]int), args.Error(1)
}

// TestPackingService_Calculate_Success tests successful calculation
func TestPackingService_Calculate_Success(t *testing.T) {
	// Arrange
	mockCalculator := new(MockCalculator)
	numberOfItems := 100
	boxCapacity := []int{10, 20, 30}
	expectedResult := map[int]int{10: 2, 20: 3, 30: 1}

	mockCalculator.On("CalculateOptimalPack", numberOfItems, boxCapacity).
		Return(expectedResult, nil).
		Once()

	service := NewPackingService(mockCalculator)
	request := model.PackRequest{
		NumberOfItems: numberOfItems,
		BoxCapacity:   boxCapacity,
	}

	// Act
	response, err := service.Calculate(request)

	// Assert
	assert.NoError(t, err, "Calculate should not return an error")
	assert.Equal(t, numberOfItems, response.TotalItems, "TotalItems should match request")
	assert.Len(t, response.Results, 3, "Results should have 3 items")

	// Verify the structure of results
	resultMap := make(map[int]int)
	for _, result := range response.Results {
		resultMap[result.Capacity] = result.BoxCount
	}
	assert.Equal(t, expectedResult, resultMap, "Results should match expected packing")

	// Verify mock was called with correct arguments
	mockCalculator.AssertExpectations(t)
	mockCalculator.AssertCalled(t, "CalculateOptimalPack", numberOfItems, boxCapacity)
}

// TestPackingService_Calculate_EmptyResult tests successful calculation with no results
func TestPackingService_Calculate_EmptyResult(t *testing.T) {
	// Arrange
	mockCalculator := new(MockCalculator)
	numberOfItems := 0
	boxCapacity := []int{10}
	expectedResult := map[int]int{}

	mockCalculator.On("CalculateOptimalPack", numberOfItems, boxCapacity).
		Return(expectedResult, nil).
		Once()

	service := NewPackingService(mockCalculator)
	request := model.PackRequest{
		NumberOfItems: numberOfItems,
		BoxCapacity:   boxCapacity,
	}

	// Act
	response, err := service.Calculate(request)

	// Assert
	assert.NoError(t, err, "Calculate should not return an error")
	assert.Equal(t, numberOfItems, response.TotalItems)
	assert.Empty(t, response.Results, "Results should be empty")
	mockCalculator.AssertExpectations(t)
}

// TestPackingService_Calculate_DuplicatePackageSizeError tests error handling for duplicate package sizes
func TestPackingService_Calculate_DuplicatePackageSizeError(t *testing.T) {
	// Arrange
	mockCalculator := new(MockCalculator)
	numberOfItems := 100
	boxCapacity := []int{10, 10, 20}
	expectedError := calculator.ErrDuplicatePackageSizes

	mockCalculator.On("CalculateOptimalPack", numberOfItems, boxCapacity).
		Return(nil, expectedError).
		Once()

	service := NewPackingService(mockCalculator)
	request := model.PackRequest{
		NumberOfItems: numberOfItems,
		BoxCapacity:   boxCapacity,
	}

	// Act
	response, err := service.Calculate(request)

	// Assert
	assert.Error(t, err, "Calculate should return an error")
	assert.Equal(t, numberOfItems, response.TotalItems, "TotalItems should still be set")
	assert.Empty(t, response.Results, "Results should be empty when error occurs")
	assert.Contains(t, err.Error(), "error calculating optimal pack", "Error message should contain context")

	mockCalculator.AssertExpectations(t)
	mockCalculator.AssertCalled(t, "CalculateOptimalPack", numberOfItems, boxCapacity)
}

// TestPackingService_Calculate_CustomError tests handling of custom errors from calculator
func TestPackingService_Calculate_CustomError(t *testing.T) {
	// Arrange
	mockCalculator := new(MockCalculator)
	numberOfItems := 100
	boxCapacity := []int{10, 20}
	customError := errors.New("custom calculator error")

	mockCalculator.On("CalculateOptimalPack", numberOfItems, boxCapacity).
		Return(nil, customError).
		Once()

	service := NewPackingService(mockCalculator)
	request := model.PackRequest{
		NumberOfItems: numberOfItems,
		BoxCapacity:   boxCapacity,
	}

	// Act
	response, err := service.Calculate(request)

	// Assert
	assert.Error(t, err, "Calculate should return an error")
	assert.Equal(t, numberOfItems, response.TotalItems)
	assert.Empty(t, response.Results)
	assert.Contains(t, err.Error(), "error calculating optimal pack")
	assert.Contains(t, err.Error(), "custom calculator error")

	mockCalculator.AssertExpectations(t)
}

// TestPackingService_Calculate_LargeNumbers tests calculation with large numbers
func TestPackingService_Calculate_LargeNumbers(t *testing.T) {
	// Arrange
	mockCalculator := new(MockCalculator)
	numberOfItems := 1000000
	boxCapacity := []int{100, 500, 1000}
	expectedResult := map[int]int{1000: 1000}

	mockCalculator.On("CalculateOptimalPack", numberOfItems, boxCapacity).
		Return(expectedResult, nil).
		Once()

	service := NewPackingService(mockCalculator)
	request := model.PackRequest{
		NumberOfItems: numberOfItems,
		BoxCapacity:   boxCapacity,
	}

	// Act
	response, err := service.Calculate(request)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, numberOfItems, response.TotalItems)
	assert.Len(t, response.Results, 1)
	assert.Equal(t, 1000, response.Results[0].Capacity)
	assert.Equal(t, 1000, response.Results[0].BoxCount)

	mockCalculator.AssertExpectations(t)
}

// TestPackingService_Calculate_VerifyArgumentsPassed tests that arguments are passed correctly to calculator
func TestPackingService_Calculate_VerifyArgumentsPassed(t *testing.T) {
	// Arrange
	mockCalculator := new(MockCalculator)
	numberOfItems := 50
	boxCapacity := []int{5, 15, 25}

	// Use Run to verify and capture arguments
	mockCalculator.On("CalculateOptimalPack", numberOfItems, boxCapacity).
		Run(func(args mock.Arguments) {
			num := args.Get(0).(int)
			sizes := args.Get(1).([]int)
			assert.Equal(t, numberOfItems, num, "NumberOfItems argument should match")
			assert.Equal(t, boxCapacity, sizes, "BoxCapacity argument should match")
		}).
		Return(map[int]int{5: 10}, nil).
		Once()

	service := NewPackingService(mockCalculator)
	request := model.PackRequest{
		NumberOfItems: numberOfItems,
		BoxCapacity:   boxCapacity,
	}

	// Act
	_, err := service.Calculate(request)

	// Assert
	assert.NoError(t, err)
	mockCalculator.AssertNumberOfCalls(t, "CalculateOptimalPack", 1)
	mockCalculator.AssertExpectations(t)
}

// TestPackingService_Calculate_SingleBoxCapacity tests calculation with single box capacity
func TestPackingService_Calculate_SingleBoxCapacity(t *testing.T) {
	// Arrange
	mockCalculator := new(MockCalculator)
	numberOfItems := 100
	boxCapacity := []int{50}
	expectedResult := map[int]int{50: 2}

	mockCalculator.On("CalculateOptimalPack", numberOfItems, boxCapacity).
		Return(expectedResult, nil).
		Once()

	service := NewPackingService(mockCalculator)
	request := model.PackRequest{
		NumberOfItems: numberOfItems,
		BoxCapacity:   boxCapacity,
	}

	// Act
	response, err := service.Calculate(request)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response.Results))
	assert.Equal(t, 50, response.Results[0].Capacity)
	assert.Equal(t, 2, response.Results[0].BoxCount)

	mockCalculator.AssertExpectations(t)
}

// TestPackingService_NewPackingService tests service initialization
func TestPackingService_NewPackingService(t *testing.T) {
	// Arrange
	mockCalculator := new(MockCalculator)

	// Act
	service := NewPackingService(mockCalculator)

	// Assert
	assert.NotNil(t, service, "Service should be created successfully")
	assert.Implements(t, (*PackingService)(nil), service, "Service should implement PackingService interface")
}
