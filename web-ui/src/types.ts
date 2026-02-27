/**
 * Represents the result of packing items into a box.
 */
export interface PackResult {
  capacity: number
  boxCount: number
}

/**
 * Response from API containing total items and the packing results.
 */
export interface CalculationResponse {
  totalItems: number
  results: PackResult[]
}
