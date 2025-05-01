package analysis

import (
	"github.com/youthtrouble/sliding-window/internal/models"
)

// Signal represents a trading signal
type Signal string

const (
	BuySignal  Signal = "Buy Signal"
	SellSignal Signal = "Sell Signal"
	Hold       Signal = "Hold"
)

// SignalsAnalyzer generates trading signals
type SignalsAnalyzer struct {
	windowSize int
}

// NewSignalsAnalyzer creates a new SignalsAnalyzer
func NewSignalsAnalyzer(windowSize int) *SignalsAnalyzer {
	return &SignalsAnalyzer{windowSize: windowSize}
}

// Generate generates trading signals
func (a *SignalsAnalyzer) Generate(prices models.StockPrices) []Signal {
	if len(prices) < a.windowSize {
		return nil
	}

	signals := make([]Signal, len(prices)-a.windowSize+1)

	for i := 0; i <= len(prices)-a.windowSize; i++ {
		window := prices.GetWindow(i, a.windowSize)
		avg := window.Average()
		currentPrice := window[len(window)-1].Price

		switch {
		case currentPrice < avg*0.95:
			signals[i] = BuySignal
		case currentPrice > avg*1.05:
			signals[i] = SellSignal
		default:
			signals[i] = Hold
		}
	}

	return signals
}