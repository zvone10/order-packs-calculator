import { useEffect } from 'react'
import { useCalculatorStore } from './store/calculatorStore'
import { PackSizesInput } from './components/PackSizesInput'
import { CalculationSection } from './components/CalculationSection'
import { ResultsTable } from './components/ResultsTable'

/**
 * Entry point of the application. Loads pack sizes from localStorage on mount and renders the main components.
 */
function App() {
  const loadPackSizes = useCalculatorStore((state) => state.loadPackSizes)

  useEffect(() => {
    loadPackSizes()
  }, [loadPackSizes])

  return (
    <div className="app">
      <header className="header">
        <h1>Order Packs Calculator</h1>
      </header>
      <main className="main">
        <PackSizesInput />
        <CalculationSection />
        <ResultsTable />
      </main>
    </div>
  )
}

export default App
