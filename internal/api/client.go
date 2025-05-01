package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dejiajibola/interview-prep/sliding-window/internal/models"
	"github.com/dejiajibola/interview-prep/sliding-window/pkg/config"
)

const timeLayout = "2006-01-02"

// Client represents an Alpha Vantage API client
type Client struct {
	config *config.Config
}

// NewClient creates a new Alpha Vantage API client
func NewClient(cfg *config.Config) *Client {
	return &Client{config: cfg}
}

// FetchStockData gets stock data from Alpha Vantage API
func (c *Client) FetchStockData() (models.StockPrices, error) {
	url := fmt.Sprintf("%s?function=TIME_SERIES_DAILY&symbol=%s&outputsize=%s&apikey=%s&datatype=%s",
		c.config.BaseURL, c.config.Symbol, c.config.OutputSize, c.config.APIKey, c.config.DataType)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var apiResponse AlphaVantageResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %v", err)
	}

	var prices models.StockPrices
	for dateStr, data := range apiResponse.TimeSeries {
		date, err := time.Parse(timeLayout, dateStr)
		if err != nil {
			continue
		}

		price, err := parseFloat(data.Close)
		if err != nil {
			continue
		}

		prices = append(prices, models.StockPrice{Date: date, Price: price})
	}

	prices.SortByDate()
	return prices, nil
}

func parseFloat(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
} 