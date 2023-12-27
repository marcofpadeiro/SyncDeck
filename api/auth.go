package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func APIKeyMiddleware(config Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKeyHeader := c.GetHeader("Authorization")

		if apiKeyHeader != config.Api_Key {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}
