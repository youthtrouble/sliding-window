package analysis_test

import (
	"testing"
	"time"

	"github.com/dejiajibola/interview-prep/sliding-window/internal/analysis"
	"github.com/dejiajibola/interview-prep/sliding-window/internal/models"
)

func TestTrendsAnalyzer(t *testing.T) {
	prices := models.StockPrices{
		{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
		{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 110.0}, // +10%
		{Date: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC), Price: 90.0},  // -18.18%
		{Date: time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC), Price: 95.0},  // +5.56%
		{Date: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC), Price: 85.0},  // -10.53%
	}

	analyzer := analysis.NewTrendsAnalyzer(3)

	t.Run("Identify", func(t *testing.T) {
		trends := analyzer.Identify(prices)
		if len(trends) != 3 {
			t.Errorf("Expected 3 trends, got %d", len(trends))
		}

		// First window: 100 -> 110 -> 90
		// Change = (90 - 100) / 100 * 100 = -10%
		if trends[0] != analysis.StrongDowntrend {
			t.Errorf("Expected first trend StrongDowntrend, got %v", trends[0])
		}

		// Second window: 110 -> 90 -> 95
		// Change = (95 - 110) / 110 * 100 ≈ -13.64%
		if trends[1] != analysis.StrongDowntrend {
			t.Errorf("Expected second trend StrongDowntrend, got %v", trends[1])
		}

		// Third window: 90 -> 95 -> 85
		// Change = (85 - 90) / 90 * 100 ≈ -5.56%
		if trends[2] != analysis.StrongDowntrend {
			t.Errorf("Expected third trend StrongDowntrend, got %v", trends[2])
		}
	})

	t.Run("InsufficientData", func(t *testing.T) {
		shortPrices := models.StockPrices{
			{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
			{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		}
		trends := analyzer.Identify(shortPrices)
		if trends != nil {
			t.Error("Expected nil for insufficient data")
		}
	})
} 