package models_test

import (
	"testing"
	"time"

	"github.com/youthtrouble/sliding-window/internal/models"
)

func TestStockPrices(t *testing.T) {
	prices := models.StockPrices{
		{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
		{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		{Date: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC), Price: 95.0},
	}

	t.Run("SortByDate", func(t *testing.T) {
		unsorted := models.StockPrices{
			{Date: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC), Price: 95.0},
			{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), Price: 100.0},
			{Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), Price: 105.0},
		}
		unsorted.SortByDate()

		if unsorted[0].Date != time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC) {
			t.Errorf("Expected first date to be 2023-01-01, got %v", unsorted[0].Date)
		}
	})

	t.Run("GetWindow", func(t *testing.T) {
		window := prices.GetWindow(0, 2)
		if len(window) != 2 {
			t.Errorf("Expected window length 2, got %d", len(window))
		}
		if window[0].Price != 100.0 || window[1].Price != 105.0 {
			t.Errorf("Expected prices [100.0, 105.0], got [%f, %f]", window[0].Price, window[1].Price)
		}
	})

	t.Run("Average", func(t *testing.T) {
		avg := prices.Average()
		expected := (100.0 + 105.0 + 95.0) / 3.0
		if avg != expected {
			t.Errorf("Expected average %f, got %f", expected, avg)
		}
	})

	t.Run("Max", func(t *testing.T) {
		max := prices.Max()
		if max != 105.0 {
			t.Errorf("Expected max 105.0, got %f", max)
		}
	})

	t.Run("Min", func(t *testing.T) {
		min := prices.Min()
		if min != 95.0 {
			t.Errorf("Expected min 95.0, got %f", min)
		}
	})
}