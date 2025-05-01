package analysis

import (
	"github.com/dejiajibola/interview-prep/sliding-window/internal/models"
)

// MovingAverageAnalyzer calculates moving averages for stock prices
type MovingAverageAnalyzer struct {
	windowSize int
}

// NewMovingAverageAnalyzer creates a new MovingAverageAnalyzer
func NewMovingAverageAnalyzer(windowSize int) *MovingAverageAnalyzer {
	return &MovingAverageAnalyzer{windowSize: windowSize}
}

// Calculate calculates moving averages for the given stock prices
func (a *MovingAverageAnalyzer) Calculate(prices models.StockPrices) []float64 {
	if len(prices) < a.windowSize {
		return nil
	}

	averages := make([]float64, len(prices)-a.windowSize+1)

	for i := 0; i <= len(prices)-a.windowSize; i++ {
		window := prices.GetWindow(i, a.windowSize)
		averages[i] = window.Average()
	}

	return averages
} 