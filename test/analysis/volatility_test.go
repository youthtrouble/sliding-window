package analysis_test

import (
	"math"
	"testing"
	"time"

	"github.com/youthtrouble/sliding-window/internal/analysis"
	"github.com/youthtrouble/sliding-window/internal/models"
)

func TestVolatilityAnalyzer(t *testing.T) {
	prices := models.StockPrices{
		{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
		{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		{Date: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC), Price: 95.0},
		{Date: time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC), Price: 110.0},
		{Date: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC), Price: 115.0},
	}

	analyzer := analysis.NewVolatilityAnalyzer(3)

	t.Run("Calculate", func(t *testing.T) {
		volatilities := analyzer.Calculate(prices)
		if len(volatilities) != 3 {
			t.Errorf("Expected 3 volatilities, got %d", len(volatilities))
		}

		// First window: [100, 105, 95]
		// Mean = 100
		// Variance = ((100-100)^2 + (105-100)^2 + (95-100)^2) / 3 = (0 + 25 + 25) / 3 ≈ 16.67
		// Std dev = sqrt(16.67) ≈ 4.08
		expectedVolatility := math.Sqrt((0 + 25 + 25) / 3.0)
		if math.Abs(volatilities[0]-expectedVolatility) > 0.01 {
			t.Errorf("Expected first volatility ≈4.08, got %f", volatilities[0])
		}
	})

	t.Run("InsufficientData", func(t *testing.T) {
		shortPrices := models.StockPrices{
			{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
			{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		}
		volatilities := analyzer.Calculate(shortPrices)
		if volatilities != nil {
			t.Error("Expected nil for insufficient data")
		}
	})
}