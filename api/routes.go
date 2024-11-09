package api

import (
	"GeO-Locator/api/handlers"
	"github.com/gin-gonic/gin"
)

func registerRoutes(engine *gin.Engine) {
	api := engine.Group("api")

	api.GET("", handlers.GetIndex)
	api.GET("index", handlers.GetIndex)

	api.GET("geo-locations", handlers.GetGeoLocation)
}
