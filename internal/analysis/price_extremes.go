package analysis

import (
	"github.com/youthtrouble/sliding-window/internal/models"
)

// PriceExtremesAnalyzer finds maximum and minimum prices in windows
type PriceExtremesAnalyzer struct {
	windowSize int
}

// NewPriceExtremesAnalyzer creates a new PriceExtremesAnalyzer
func NewPriceExtremesAnalyzer(windowSize int) *PriceExtremesAnalyzer {
	return &PriceExtremesAnalyzer{windowSize: windowSize}
}

// FindMaxPrices finds maximum prices in windows
func (a *PriceExtremesAnalyzer) FindMaxPrices(prices models.StockPrices) []float64 {
	if len(prices) < a.windowSize {
		return nil
	}

	maxPrices := make([]float64, len(prices)-a.windowSize+1)

	for i := 0; i <= len(prices)-a.windowSize; i++ {
		window := prices.GetWindow(i, a.windowSize)
		maxPrices[i] = window.Max()
	}

	return maxPrices
}

// FindMinPrices finds minimum prices in windows
func (a *PriceExtremesAnalyzer) FindMinPrices(prices models.StockPrices) []float64 {
	if len(prices) < a.windowSize {
		return nil
	}

	minPrices := make([]float64, len(prices)-a.windowSize+1)

	for i := 0; i <= len(prices)-a.windowSize; i++ {
		window := prices.GetWindow(i, a.windowSize)
		minPrices[i] = window.Min()
	}

	return minPrices
}