// service/packing_service.go
package service

import (
	"order-pack-calculator-api/internal/calculator"
	"order-pack-calculator-api/internal/model"
)

type PackingService interface {
	Calculate(req model.PackRequest) model.PackResponse
}

type packingService struct {
	packagingCalculator calculator.Calculator
}

func NewPackingService(calculator calculator.Calculator) PackingService {
	return &packingService{
		packagingCalculator: calculator,
	}
}

func (s *packingService) Calculate(req model.PackRequest) model.PackResponse {
	results := make([]model.PackResult, 0, len(req.BoxCapacity))

	optimalPack, err := s.packagingCalculator.CalculateOptimalPack(req.NumberOfItems, req.BoxCapacity)
	if err != nil {
		return model.PackResponse{
			TotalItems: req.NumberOfItems,
			Results:    results,
		}
	}

	for capcity, numberOfBoxes := range optimalPack {
		results = append(results, model.PackResult{
			Capacity: capcity,
			BoxCount: numberOfBoxes,
		})
	}

	return model.PackResponse{
		TotalItems: req.NumberOfItems,
		Results:    results,
	}
}
