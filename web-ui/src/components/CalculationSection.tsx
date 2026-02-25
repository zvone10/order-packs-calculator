import { useCalculatorStore } from '../store/calculatorStore'

export function CalculationSection() {
  const { numberOfItems, setNumberOfItems, calculate, loading, error } =
    useCalculatorStore()

  const handleCalculate = async () => {
    await calculate()
  }

  return (
    <section className="section">
      <h2>Calculate</h2>
      <div className="input-group">
        <input
          type="number"
          value={numberOfItems}
          onChange={(e) => setNumberOfItems(Number(e.target.value))}
          placeholder="Number of items"
          min="1"
          disabled={loading}
        />
        <button
          onClick={handleCalculate}
          disabled={loading}
          className="btn-primary"
        >
          {loading ? 'Calculating...' : 'Calculate'}
        </button>
      </div>
      {error && <p className="error">{error}</p>}
    </section>
  )
}
