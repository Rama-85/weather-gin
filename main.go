package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type weather struct {
	ID             string  `json:"id"`
	DeviceName     string  `json:"devicename"`
	AirTemperature float64 `json:"airtemperature"`
	Humidity       float64 `json:"humidity"`
	Windspeed      float64 `json:"windspeed"`
}

var wt = []weather{
	{ID: "1", DeviceName: "Thermometer", AirTemperature: 22.4, Humidity: 56.99, Windspeed: 10.5},
	{ID: "2", DeviceName: "Anemometer", AirTemperature: 33.6, Humidity: 36.99, Windspeed: 21.6},
	{ID: "3", DeviceName: "Barometer", AirTemperature: 42.3, Humidity: 16.99, Windspeed: 8.5},
}

func main() {
	router := gin.Default()
	router.GET("/wt", getWeather)
	router.GET("/wt/:id", getWeatherByID)
	router.POST("/wt/", postWeather)
	router.DELETE("/wt/:id", DeleteWeather)

	router.Run("localhost:8081")
}

func getWeather(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, wt)
}

func postWeather(c *gin.Context) {
	var newWeather weather

	if err := c.BindJSON(&newWeather); err != nil {
		return
	}

	wt = append(wt, newWeather)
	c.IndentedJSON(http.StatusCreated, newWeather)
}

func getWeatherByID(c *gin.Context) {
	id := c.Param("id")

	for _, w := range wt {
		if w.ID == id {
			c.IndentedJSON(http.StatusOK, w)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "weather not found"})
}

func DeleteWeather(c *gin.Context) {

	id := c.Param("id")
	for i, w := range wt {
		if w.ID == id {
			wt = append(wt[:i], wt[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})

}
