package analysis_test

import (
	"testing"
	"time"

	"github.com/dejiajibola/interview-prep/sliding-window/internal/models"
	"github.com/youthtrouble/sliding-window/internal/analysis"
)

func TestSignalsAnalyzer(t *testing.T) {
	prices := models.StockPrices{
		{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
		{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		{Date: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC), Price: 95.0},
		{Date: time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC), Price: 110.0},
		{Date: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC), Price: 115.0},
	}

	analyzer := analysis.NewSignalsAnalyzer(3)

	t.Run("Generate", func(t *testing.T) {
		signals := analyzer.Generate(prices)
		if len(signals) != 3 {
			t.Errorf("Expected 3 signals, got %d", len(signals))
		}

		// First window: [100, 105, 95]
		// Average = 100
		// Current price = 95 (5% below average)
		if signals[0] != analysis.BuySignal {
			t.Errorf("Expected first signal BuySignal, got %v", signals[0])
		}

		// Second window: [105, 95, 110]
		// Average ≈ 103.33
		// Current price = 110 (above average but not 5% above)
		if signals[1] != analysis.Hold {
			t.Errorf("Expected second signal Hold, got %v", signals[1])
		}

		// Third window: [95, 110, 115]
		// Average ≈ 106.67
		// Current price = 115 (above 5% above average)
		if signals[2] != analysis.SellSignal {
			t.Errorf("Expected third signal SellSignal, got %v", signals[2])
		}
	})

	t.Run("InsufficientData", func(t *testing.T) {
		shortPrices := models.StockPrices{
			{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
			{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		}
		signals := analyzer.Generate(shortPrices)
		if signals != nil {
			t.Error("Expected nil for insufficient data")
		}
	})
}