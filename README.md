# Weather App

The Weather App is a Go-based application that provides weather information for a given city. It uses a third-party weather API (https://www.weatherapi.com/) to fetch real-time weather data and caches the results in Redis to improve performance and reduce API calls.

## Features

- Fetches real-time weather data for a specified city.
- Caches weather data in Redis with an expiration time to reduce redundant API calls.
- Provides a RESTful API with endpoints for:
  - Health check.
  - Fetching weather data for a city.
- Implements a timeout mechanism for API calls to ensure responsiveness.
- Normalizes city names to handle variations in input (e.g., "New-York" vs. "New York").
- Configurable via environment variables.

---

## Project Structure

```
weather-app/
├── cmd/
│   ├── api.go          # Main application logic
│   ├── handler.go      # HTTP handlers for API endpoints
│   ├── main.go         # Entry point of the application
├── internal/
│   ├── env/
│   │   └── env.go      # Utility functions for environment variable management
│   ├── model/
│   │   └── model.go    # Data models for weather API responses
│   ├── repository/
│   │   └── repository.go # Redis repository for caching weather data
│   ├── store/
│   │   └── redis.go    # Redis client setup
├── service/
│   └── service.go      # Business logic for fetching and caching weather data
├── types/
│   └── types.go        # Shared types and data structures
├── utils/
│   └── utils.go        # Utility functions (e.g., API call helper)
```

---

## Endpoints

### Health Check
- **URL**: `/health`
- **Method**: `GET`
- **Description**: Returns the status of the application.
- **Response**:
  ```json
  {
    "status": "ok",
    "env": "8080",
    "version": "1.1.0"
  }
  ```

### Get Weather by City
- **URL**: `/v1/weather`
- **Method**: `GET`
- **Query Parameters**:
  - `city` (required): The name of the city to fetch weather data for.
- **Description**: Fetches weather data for the specified city. If the data is cached in Redis, it retrieves it from there; otherwise, it fetches it from the weather API and stores it in Redis.
- **Response**:
  ```json
  {
    "status": "success",
    "data": {
      "name": "New York",
      "region": "New York",
      "country": "United States",
      "latitude": 40.7128,
      "longitude": -74.006,
      "localtime": "2025-05-06 10:00",
      "temp_c": 20.5,
      "temp_f": 68.9,
      "last_updated": "2025-05-06 09:45",
      "condition": {
        "text": "Partly cloudy",
        "icon": "//cdn.weatherapi.com/weather/64x64/day/116.png",
        "code": 1003
      },
      "uv": 5.0
    }
  }
  ```

---

## Setup Instructions

### Prerequisites
- Go 1.18 or later
- Redis server
- A weather API key (e.g., from WeatherAPI)

### Environment Variables
The application uses the following environment variables:

| Variable           | Default Value       | Description                              |
|--------------------|---------------------|------------------------------------------|
| `ADDR`             | `:8080`            | Address and port for the server          |
| `WEATHER_API_URL`  | (required)          | Base URL of the weather API              |
| `WEATHER_API_KEY`  | (required)          | API key for the weather API              |
| `REDIS_ADDR`       | `localhost:6379`    | Redis server address                     |
| `REDIS_PW`         | `""`               | Redis password (if any)                  |
| `REDIS_DB`         | `0`                | Redis database index                     |
| `TIME_DURATION`    | `10`               | Timeout duration for API calls (seconds) |

### Steps to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/AyesoroFemi/weather-api.git
   cd weather-app
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   - Use `direnv` or create a `.env` file with the required variables.

4. Start the Redis server:
   ```bash
   redis-server
   ```

5. Run the application:
   ```bash
   go run cmd/main.go
   ```

6. Access the application:
   - Health check: `http://localhost:8080/health`
   - Weather endpoint: `http://localhost:8080/v1/weather?city=New-York`

---

## Key Components

### Redis Caching
- Weather data is cached in Redis with a configurable expiration time (default: 10 minutes).
- Reduces redundant API calls and improves performance.

### City Name Normalization
- Handles variations in city name input (e.g., "New-York" vs. "New York").
- Uses a utility function to normalize city names for consistent comparison.

### Timeout Mechanism
- Ensures API calls and Redis operations do not hang indefinitely.
- Configurable via the `TIME_DURATION` environment variable.

---

## Future Enhancements
- Add support for multiple weather APIs with fallback mechanisms.
- Implement rate limiting to prevent abuse of the API.
- Add unit tests for all components.
- Improve error handling and logging.

---

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push them to your fork.
4. Submit a pull request.

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

---