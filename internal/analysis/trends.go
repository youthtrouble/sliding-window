package analysis

import (
	"github.com/youthtrouble/sliding-window/internal/models"
)

// Trend represents a price trend
type Trend string

const (
	StrongUptrend   Trend = "Strong Uptrend"
	Uptrend         Trend = "Uptrend"
	StrongDowntrend Trend = "Strong Downtrend"
	Downtrend       Trend = "Downtrend"
	Sideways        Trend = "Sideways"
)

// TrendsAnalyzer identifies price trends in windows
type TrendsAnalyzer struct {
	windowSize int
}

// NewTrendsAnalyzer creates a new TrendsAnalyzer
func NewTrendsAnalyzer(windowSize int) *TrendsAnalyzer {
	return &TrendsAnalyzer{windowSize: windowSize}
}

// Identify identifies price trends in windows
func (a *TrendsAnalyzer) Identify(prices models.StockPrices) []Trend {
	if len(prices) < a.windowSize {
		return nil
	}

	trends := make([]Trend, len(prices)-a.windowSize+1)

	for i := 0; i <= len(prices)-a.windowSize; i++ {
		window := prices.GetWindow(i, a.windowSize)
		startPrice := window[0].Price
		endPrice := window[len(window)-1].Price
		change := (endPrice - startPrice) / startPrice * 100

		switch {
		case change > 5:
			trends[i] = StrongUptrend
		case change > 1:
			trends[i] = Uptrend
		case change < -5:
			trends[i] = StrongDowntrend
		case change < -1:
			trends[i] = Downtrend
		default:
			trends[i] = Sideways
		}
	}

	return trends
}