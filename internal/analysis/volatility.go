package analysis

import (
	"math"

	"github.com/youthtrouble/sliding-window/internal/models"
)

// VolatilityAnalyzer calculates price volatility in windows
type VolatilityAnalyzer struct {
	windowSize int
}

// NewVolatilityAnalyzer creates a new VolatilityAnalyzer
func NewVolatilityAnalyzer(windowSize int) *VolatilityAnalyzer {
	return &VolatilityAnalyzer{windowSize: windowSize}
}

// Calculate calculates price volatility in windows
func (a *VolatilityAnalyzer) Calculate(prices models.StockPrices) []float64 {
	if len(prices) < a.windowSize {
		return nil
	}

	volatilities := make([]float64, len(prices)-a.windowSize+1)

	for i := 0; i <= len(prices)-a.windowSize; i++ {
		window := prices.GetWindow(i, a.windowSize)
		avg := window.Average()

		variance := 0.0
		for _, price := range window {
			diff := price.Price - avg
			variance += diff * diff
		}
		variance /= float64(a.windowSize)

		volatilities[i] = math.Sqrt(variance)
	}

	return volatilities
}