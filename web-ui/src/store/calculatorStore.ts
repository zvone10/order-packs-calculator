import { create } from 'zustand'
import { CalculationResponse } from '../types'
import { joinUrl } from '../utils/urlUtils'

const STORAGE_KEY = 'packSizes'

/** Store of app state */
interface CalculatorStore {
  packSizes: number[]
  numberOfItems: number
  results: CalculationResponse | null
  loading: boolean
  error: string | null

  // Pack sizes actions
  addPackSize: () => void
  removePackSize: (index: number) => void
  updatePackSize: (index: number, value: number) => void
  loadPackSizes: () => void
  clearPackSizes: () => void

  // Calculation actions
  setNumberOfItems: (count: number) => void
  setResults: (data: CalculationResponse) => void
  setLoading: (loading: boolean) => void
  setError: (error: string | null) => void
  calculate: () => Promise<void>
}

export const useCalculatorStore = create<CalculatorStore>((set, get) => ({
  packSizes: [],
  numberOfItems: 0,
  results: null,
  loading: false,
  error: null,

  addPackSize: () => {
    set((state) => {
      const updated = [...state.packSizes, 0]
      localStorage.setItem(STORAGE_KEY, JSON.stringify(updated))
      return { packSizes: updated }
    })
  },

  removePackSize: (index: number) => {
    set((state) => {
      const updated = state.packSizes.filter((_, i) => i !== index)
      localStorage.setItem(STORAGE_KEY, JSON.stringify(updated))
      return { packSizes: updated }
    })
  },

  updatePackSize: (index: number, value: number) => {
    set((state) => {
      const updated = [...state.packSizes]
      updated[index] = value
      localStorage.setItem(STORAGE_KEY, JSON.stringify(updated))
      return { packSizes: updated }
    })
  },

  loadPackSizes: () => {
    const stored = localStorage.getItem(STORAGE_KEY)
    if (stored) {
      try {
        set({ packSizes: JSON.parse(stored) })
      } catch {
        set({ packSizes: [] })
      }
    }
  },

  clearPackSizes: () => {
    localStorage.removeItem(STORAGE_KEY)
    set({ packSizes: [] })
  },

  setNumberOfItems: (count: number) => {
    set({ numberOfItems: count })
  },

  setResults: (data: CalculationResponse) => {
    set({ results: data })
  },

  setLoading: (loading: boolean) => {
    set({ loading })
  },

  setError: (error: string | null) => {
    set({ error })
  },

  calculate: async () => {
    const state = get()
    if (state.packSizes.length === 0) {
      set({ error: 'Please add at least one pack size' })
      return
    }
    if (state.numberOfItems <= 0) {
      set({ error: 'Please enter a valid number of items' })
      return
    }

    set({ loading: true, error: null })
    try {
      const apiUrl = joinUrl(import.meta.env.VITE_API_URL, 'api/v1/pack-calculation')
      const response = await fetch(apiUrl, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          numberOfItems: state.numberOfItems,
          boxCapacity: state.packSizes,
        }),
      })

      if (!response.ok) {
        throw new Error(`HTTP error status: ${response.status}`)
      }

      const data: CalculationResponse = await response.json()
      set({ results: data, error: null })
    } catch (err) {
      set({ error: err instanceof Error ? err.message : 'An error occurred' })
    } finally {
      set({ loading: false })
    }
  },
}))
