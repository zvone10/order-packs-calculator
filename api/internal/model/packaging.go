package model

type PackRequest struct {
	NumberOfItems int   `json:"numberOfItems" binding:"required,gt=0"`
	BoxCapacity   []int `json:"boxCapacity"   binding:"required,min=1,dive,gt=0"`
}

type PackResult struct {
	Capacity int `json:"capacity"`
	BoxCount int `json:"boxCount"`
}

/**
* PackResponse represents the response structure for the packing calculation.
 */
type PackResponse struct {
	TotalItems int          `json:"totalItems"`
	Results    []PackResult `json:"results"`
}
