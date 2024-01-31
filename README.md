# GO Weather App

This is Go application that fetches weather information from the OpenWeatherMap API based on user input and returns the weather data in JSON format.

## Table of Contents

- [Features](#features)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)

## Features

- Fetch weather data for a specific city.
- Return weather data in JSON format.
- Error handling for missing API keys and invalid requests.

## Environment Variables

The following environment variables are used in the application:

- OPEN_WEATHER_API_KEY: API key for accessing the OpenWeatherMap API. Get your API key [here](https://home.openweathermap.org/api_keys)

Ensure these variables are set before running the application.

## Installation

To install the Weather App, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/coolpythoncodes/weather-go-api.git
```

2. Navigate to the project directory:

```bash
cd weather-go-api
```

3. Set up environment variables by creating a .env file in the root directory and adding the necessary variables (see [Environment Variables](#environment-variables)). Note: this is applicable to those using docker compose to build the app.

## Usage

To run the Weather App(not applicable to people using docker compose),

Set your environment variable by running the following commands on your terminal

```bash
export OPEN_WEATHER_API_KEY="insert your api key here"
```

```bash
go run main.go
```

For Docker fans

```bash
docker compose up --build
```

## Endpoints

The Weather App provides the following endpoints:

GET /weather/:city: Fetches weather data for the specified city.

Example:

```bash
curl http://localhost:8080/weather/London
```

Response:

```json
{
	"error": false,
	"data": {
		"name": "London",
		"weather": [
			{
				"main": "Clouds",
				"description": "scattered clouds"
			}
		],
		"main": {
			"temp": 12.34
		}
	}
}
```
