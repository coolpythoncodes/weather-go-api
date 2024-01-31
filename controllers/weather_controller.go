package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coolpythoncodes/weather-api/models"
	"github.com/gin-gonic/gin"
)

func GetWeatherDetails(c *gin.Context) {

	openWeatherApiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if openWeatherApiKey == "" {
		log.Fatal("OPEN_WEATHER_API_KEY is not provided in the environment file. Please set a valid API key.")
		return
	}

	weatherRequest := models.WeatherRequest{}
	if err := c.ShouldBindUri(&weatherRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + weatherRequest.City + "&appid=" + openWeatherApiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	defer resp.Body.Close()

	weatherResponseData := models.WeatherResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&weatherResponseData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	}
	fmt.Println("weatherResponseData", weatherResponseData)
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  weatherResponseData,
	})

}
