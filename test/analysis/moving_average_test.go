package analysis_test

import (
	"testing"
	"time"

	"github.com/youthtrouble/sliding-window/internal/analysis"
	"github.com/youthtrouble/sliding-window/internal/models"
)

func TestMovingAverageAnalyzer(t *testing.T) {
	prices := models.StockPrices{
		{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
		{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		{Date: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC), Price: 95.0},
		{Date: time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC), Price: 110.0},
		{Date: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC), Price: 115.0},
	}

	analyzer := analysis.NewMovingAverageAnalyzer(3)

	t.Run("Calculate", func(t *testing.T) {
		averages := analyzer.Calculate(prices)
		if len(averages) != 3 {
			t.Errorf("Expected 3 averages, got %d", len(averages))
		}

		// First window: (100 + 105 + 95) / 3 = 100
		if averages[0] != 100.0 {
			t.Errorf("Expected first average 100.0, got %f", averages[0])
		}

		// Second window: (105 + 95 + 110) / 3 ≈ 103.33
		if averages[1] != (105.0+95.0+110.0)/3.0 {
			t.Errorf("Expected second average ≈103.33, got %f", averages[1])
		}

		// Third window: (95 + 110 + 115) / 3 ≈ 106.67
		if averages[2] != (95.0+110.0+115.0)/3.0 {
			t.Errorf("Expected third average ≈106.67, got %f", averages[2])
		}
	})

	t.Run("InsufficientData", func(t *testing.T) {
		shortPrices := models.StockPrices{
			{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
			{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		}
		averages := analyzer.Calculate(shortPrices)
		if averages != nil {
			t.Error("Expected nil for insufficient data")
		}
	})
}