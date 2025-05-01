package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/youthtrouble/sliding-window/internal/api"
	"github.com/youthtrouble/sliding-window/pkg/config"
)

func TestClient_FetchStockData(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request parameters
		if r.URL.Query().Get("function") != "TIME_SERIES_DAILY" {
			t.Errorf("Expected function=TIME_SERIES_DAILY, got %s", r.URL.Query().Get("function"))
		}
		if r.URL.Query().Get("symbol") != "MSFT" {
			t.Errorf("Expected symbol=MSFT, got %s", r.URL.Query().Get("symbol"))
		}

		// Create a mock response
		response := api.AlphaVantageResponse{
			TimeSeries: map[string]struct {
				Close string `json:"4. close"`
			}{
				"2023-01-01": {"4. close": "100.0"},
				"2023-01-02": {"4. close": "105.0"},
				"2023-01-03": {"4. close": "95.0"},
			},
		}

		// Write the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create a test config
	cfg := &config.Config{
		BaseURL:    server.URL,
		Symbol:     "MSFT",
		OutputSize: "compact",
		DataType:   "json",
		APIKey:     "test-key",
	}

	// Create a client
	client := api.NewClient(cfg)

	// Test fetching stock data
	prices, err := client.FetchStockData()
	if err != nil {
		t.Fatalf("Error fetching stock data: %v", err)
	}

	// Verify the results
	if len(prices) != 3 {
		t.Errorf("Expected 3 prices, got %d", len(prices))
	}

	// Verify the first price
	if prices[0].Date != time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Expected first date 2023-01-01, got %v", prices[0].Date)
	}
	if prices[0].Price != 100.0 {
		t.Errorf("Expected first price 100.0, got %f", prices[0].Price)
	}

	// Verify the last price
	if prices[2].Date != time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC) {
		t.Errorf("Expected last date 2023-01-03, got %v", prices[2].Date)
	}
	if prices[2].Price != 95.0 {
		t.Errorf("Expected last price 95.0, got %f", prices[2].Price)
	}
}