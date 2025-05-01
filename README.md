# Stock Price Analysis with Sliding Window

A Go application that analyzes stock prices using sliding window techniques to calculate various metrics and generate trading signals.

## Features

- Fetches stock data from Alpha Vantage API
- Calculates moving averages
- Identifies price extremes (max/min)
- Measures price volatility
- Detects price trends
- Generates trading signals
- Configurable window sizes
- Comprehensive test coverage

## Project Structure

```
sliding-window/
├── cmd/
│   └── main.go              # Main application entry point
├── internal/
│   ├── api/
│   │   ├── client.go        # Alpha Vantage API client
│   │   └── types.go         # API response types
│   ├── analysis/
│   │   ├── moving_average.go # Moving average calculations
│   │   ├── price_extremes.go # Price min/max calculations
│   │   ├── volatility.go    # Volatility calculations
│   │   ├── trends.go        # Trend detection
│   │   └── signals.go       # Trading signal generation
│   └── models/
│       └── stock.go         # Stock price data models
├── pkg/
│   └── config/
│       └── config.go        # Configuration management
└── test/
    ├── api/
    │   └── client_test.go   # API client tests
    └── analysis/
        ├── moving_average_test.go
        ├── price_extremes_test.go
        ├── volatility_test.go
        ├── trends_test.go
        └── signals_test.go
```

## Prerequisites

- Go 1.16 or higher
- Alpha Vantage API key (free tier available)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/sliding-window.git
cd sliding-window
```

2. Set up your Alpha Vantage API key:
```bash
export ALPHA_VANTAGE_API_KEY="your_api_key_here"
```

## Configuration

The application can be configured through environment variables or by modifying the default values in `pkg/config/config.go`:

| Variable | Description | Default |
|----------|-------------|---------|
| ALPHA_VANTAGE_API_KEY | Alpha Vantage API key | "YOUR_API_KEY" |
| SYMBOL | Stock symbol to analyze | "MSFT" |
| WINDOW_SIZE | Size of the analysis window | 5 |
| OUTPUT_SIZE | API response size ("compact" or "full") | "compact" |
| DATA_TYPE | API response format ("json" or "csv") | "json" |

## Usage

### Running the Application

```bash
cd cmd
go run main.go
```

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o stock-analyzer ./cmd
```

## Analysis Components

### Moving Average
Calculates the average price over a sliding window of days. Useful for identifying price trends and smoothing out short-term fluctuations.

### Price Extremes
Identifies the highest and lowest prices within each window. Helps in understanding price ranges and potential support/resistance levels.

### Volatility
Measures price volatility using standard deviation. Higher volatility indicates more risk and potential trading opportunities.

### Trends
Identifies price trends based on percentage changes:
- Strong Uptrend: >5% increase
- Uptrend: 1-5% increase
- Sideways: -1% to 1% change
- Downtrend: -5% to -1% decrease
- Strong Downtrend: >5% decrease

### Trading Signals
Generates simple trading signals based on price relative to moving average:
- Buy Signal: Price drops below 95% of moving average
- Sell Signal: Price rises above 105% of moving average
- Hold: Price within 5% of moving average

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and style guidelines
- Write tests for new features
- Update documentation as needed
- Use meaningful commit messages
- Keep the code modular and maintainable

### Testing Guidelines

- Write unit tests for all new features
- Use table-driven tests where appropriate
- Test edge cases and error conditions
- Maintain high test coverage
- Run all tests before submitting PRs

## Error Handling

The application handles various error conditions:
- API connection issues
- Invalid API responses
- Insufficient data points
- Invalid price data
- Configuration errors

## Performance Considerations

- The sliding window approach is O(n) for most operations
- Memory usage is proportional to the number of data points
- API calls are rate-limited by Alpha Vantage
- Consider caching results for frequently accessed data

## Security

- API keys should never be committed to the repository
- Use environment variables for sensitive data
- Validate all input data
- Handle API responses securely

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Alpha Vantage for providing the stock data API
- The Go community for excellent tools and libraries
- Contributors who help improve the project

## Support

For support, please:
1. Check the documentation
2. Search existing issues
3. Create a new issue if needed

## Roadmap

- [ ] Add more technical indicators
- [ ] Support for multiple stocks
- [ ] Real-time data updates
- [ ] Web interface
- [ ] Historical backtesting
- [ ] Performance optimizations
- [ ] More configuration options 