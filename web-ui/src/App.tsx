import { useEffect } from 'react'
import { useCalculatorStore } from './store/calculatorStore'
import { PackSizesInput } from './components/PackSizesInput'
import { CalculationSection } from './components/CalculationSection'
import { ResultsTable } from './components/ResultsTable'

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
