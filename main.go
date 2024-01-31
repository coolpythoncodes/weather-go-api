package main

import (
	"log"
	"os"

	"github.com/coolpythoncodes/weather-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not provided in the environment file. Please set a  port.")
	}

	r := gin.Default()

	r.GET("/weather/:city", controllers.GetWeatherDetails)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(port)
}
