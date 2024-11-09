package handlers

import (
	"GeO-Locator/pkg/geoLocationProvider"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGeoLocation(c *gin.Context) {
	ip := c.ClientIP()

	provider := geoLocationProvider.NewGeoLocationProvider()
	geoLocation := provider.GetGeoLocation(ip)

	response := GetGeoLocationResponse{
		Country: geoLocation.Country,
		City:    geoLocation.City,
		Ip:      ip,
	}
	c.JSON(http.StatusOK, response)
}

type GetGeoLocationResponse struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Ip      string `json:"ip"`
}
