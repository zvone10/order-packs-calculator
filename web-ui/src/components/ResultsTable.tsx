import { useCalculatorStore } from '../store/calculatorStore'

export function ResultsTable() {
  const { results } = useCalculatorStore()

  if (!results) return null

  return (
    <section className="section">
      <h2>Results</h2>
      <p className="total">Total Items: {results.totalItems}</p>
      <table className="results-table">
        <thead>
          <tr>
            <th>Capacity</th>
            <th>Box Count</th>
          </tr>
        </thead>
        <tbody>
          {results.results.map((result, index) => (
            <tr key={index}>
              <td>{result.capacity}</td>
              <td>{result.boxCount}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </section>
  )
}
