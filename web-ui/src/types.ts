export interface PackResult {
  capacity: number
  boxCount: number
}

export interface CalculationResponse {
  totalItems: number
  results: PackResult[]
}
