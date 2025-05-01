package api

// AlphaVantageResponse represents the JSON response from Alpha Vantage
type AlphaVantageResponse struct {
	TimeSeries map[string]struct {
		Close string `json:"4. close"`
	} `json:"Time Series (Daily)"`
} 