package models

type WeatherRequest struct {
	City string `uri:"city" binding:"required"`
}
