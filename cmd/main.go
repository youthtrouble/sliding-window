package main

import (
	"fmt"
	"github.com/youthtrouble/sliding-window/internal/analysis"
	"github.com/youthtrouble/sliding-window/internal/api"
	"github.com/youthtrouble/sliding-window/pkg/config"
	"log"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Create API client
	client := api.NewClient(cfg)

	// Fetch stock data
	fmt.Printf("Fetching stock data for %s...\n", cfg.Symbol)
	prices, err := client.FetchStockData()
	if err != nil {
		log.Fatalf("Error fetching stock data: %v", err)
	}

	if len(prices) < cfg.WindowSize {
		log.Fatalf("Not enough data points (%d) for %d-day analysis", len(prices), cfg.WindowSize)
	}

	// Print the original stock price data
	fmt.Printf("\nStock Price Data for %s (last %d days):\n", cfg.Symbol, len(prices))
	fmt.Println("====================================")
	for _, price := range prices {
		fmt.Printf("%s: $%.2f\n", price.Date.Format("2006-01-02"), price.Price)
	}
	fmt.Println()

	// Create analyzers
	movingAvgAnalyzer := analysis.NewMovingAverageAnalyzer(cfg.WindowSize)
	priceExtremesAnalyzer := analysis.NewPriceExtremesAnalyzer(cfg.WindowSize)
	volatilityAnalyzer := analysis.NewVolatilityAnalyzer(cfg.WindowSize)
	trendsAnalyzer := analysis.NewTrendsAnalyzer(cfg.WindowSize)
	signalsAnalyzer := analysis.NewSignalsAnalyzer(cfg.WindowSize)

	// Calculate metrics
	averages := movingAvgAnalyzer.Calculate(prices)
	maxPrices := priceExtremesAnalyzer.FindMaxPrices(prices)
	minPrices := priceExtremesAnalyzer.FindMinPrices(prices)
	volatilities := volatilityAnalyzer.Calculate(prices)
	trends := trendsAnalyzer.Identify(prices)
	signals := signalsAnalyzer.Generate(prices)

	// Print analysis results
	fmt.Printf("%d-Day Sliding Window Analysis:\n", cfg.WindowSize)
	fmt.Println("=============================")
	for i := 0; i < len(averages); i++ {
		windowStart := prices[i].Date.Format("2006-01-02")
		windowEnd := prices[i+cfg.WindowSize-1].Date.Format("2006-01-02")

		fmt.Printf("Window %d (%s to %s):\n", i+1, windowStart, windowEnd)
		fmt.Printf("  Average Price: $%.2f\n", averages[i])
		fmt.Printf("  Max Price: $%.2f\n", maxPrices[i])
		fmt.Printf("  Min Price: $%.2f\n", minPrices[i])
		fmt.Printf("  Volatility: $%.2f\n", volatilities[i])
		fmt.Printf("  Trend: %s\n", trends[i])
		fmt.Printf("  Signal: %s\n", signals[i])
		fmt.Println()
	}
}