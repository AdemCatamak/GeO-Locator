package middleware

import (
	"GeO-Locator/internal/config"
	"GeO-Locator/internal/customErrors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Recovery Global exception handler middleware
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			configManager := config.GetConfigManager()
			showUnexpectedErrors := configManager.GetBool(config.ShowUnexpectedErrorsKey)

			if r := recover(); r != nil {
				switch err := r.(type) {

				case *customErrors.BasicCustomError:
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Message,
					})
				case *customErrors.CodedCustomError:
					c.JSON(err.Code, gin.H{
						"error": err.Message,
					})
				default:
					errorMsg := "Unexpected error occurs"
					if showUnexpectedErrors {
						errorMsg = fmt.Sprintf("String Error: %s", err)
					}

					c.JSON(http.StatusInternalServerError, gin.H{
						"error": errorMsg,
					})

					log.Panicf("unexpected error : %s", err)
				}
			}
		}()
		c.Next() // Proceed to the next handler
	}
}
