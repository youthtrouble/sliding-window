package models

import (
	"time"
)

// StockPrice represents a single day's stock price
type StockPrice struct {
	Date  time.Time
	Price float64
}

// StockPrices is a slice of StockPrice with helper methods
type StockPrices []StockPrice

// SortByDate sorts the stock prices by date in ascending order
func (p StockPrices) SortByDate() {
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			if p[i].Date.After(p[j].Date) {
				p[i], p[j] = p[j], p[i]
			}
		}
	}
}

// GetWindow returns a subslice of prices for the given window
func (p StockPrices) GetWindow(start, size int) StockPrices {
	if start < 0 || start+size > len(p) {
		return nil
	}
	return p[start : start+size]
}

// Average returns the average price of the stock prices
func (p StockPrices) Average() float64 {
	if len(p) == 0 {
		return 0
	}
	sum := 0.0
	for _, price := range p {
		sum += price.Price
	}
	return sum / float64(len(p))
}

// Max returns the maximum price in the stock prices
func (p StockPrices) Max() float64 {
	if len(p) == 0 {
		return 0
	}
	max := p[0].Price
	for _, price := range p {
		if price.Price > max {
			max = price.Price
		}
	}
	return max
}

// Min returns the minimum price in the stock prices
func (p StockPrices) Min() float64 {
	if len(p) == 0 {
		return 0
	}
	min := p[0].Price
	for _, price := range p {
		if price.Price < min {
			min = price.Price
		}
	}
	return min
} 