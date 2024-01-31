package main

import (
	"github.com/coolpythoncodes/weather-api/controllers"
	"github.com/gin-gonic/gin"
)

const Port = ":8080"

func main() {

	r := gin.Default()

	r.GET("/weather/:city", controllers.GetWeatherDetails)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(Port)
}
