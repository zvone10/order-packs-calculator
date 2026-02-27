import { useCalculatorStore } from '../store/calculatorStore'

export function PackSizesInput() {
  const {
    packSizes,
    addPackSize,
    removePackSize,
    updatePackSize,
    clearPackSizes,
  } = useCalculatorStore()

  return (
    <section className="section">
      <h2>Pack Sizes</h2>
      <div className="pack-sizes-list">
        {packSizes.map((size, index) => (
          <div key={index} className="pack-size-item">
            <input
              type="number"
              value={size}
              onChange={(e) => updatePackSize(index, Number(e.target.value))}
              placeholder="Enter pack size"
              min="1"
            />
            <button
              onClick={() => removePackSize(index)}
              className="btn-remove"
            >
              Remove
            </button>
          </div>
        ))}
      </div>
      <div className="button-group">
        <button onClick={() => addPackSize()} className="btn-primary">
          Add Pack Size
        </button>
        {packSizes.length > 0 && (
          <button onClick={clearPackSizes} className="btn-secondary">
            Clear All
          </button>
        )}
      </div>
    </section>
  )
}
