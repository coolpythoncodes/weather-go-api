package models

type MainWeatherData struct {
	Temperature float64 `json:"temp"`
}

type WeatherConditions struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type WeatherResponse struct {
	Name    string              `json:"name"`
	Weather []WeatherConditions `json:"weather"`
	Main    MainWeatherData     `json:"main"`
}
