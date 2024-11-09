package api

import (
	"GeO-Locator/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Recovery())

	registerRoutes(engine)

	return engine
}
