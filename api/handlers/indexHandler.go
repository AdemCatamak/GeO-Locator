package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func GetIndex(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	c.JSON(http.StatusOK, gin.H{
		"hostname": hostname,
		"now":      time.Now().UTC(),
	})
}
