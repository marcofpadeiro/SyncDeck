package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcofpadeiro/SyncDeck_Server/helpers"
)

func main() {
	config, err := helpers.ReadConfig()
	if err != nil {
		log.Panic("Error reading config")
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	})

	router.GET("/version/:game_id", getVersion)
	router.GET("/units", getUnits)

	router.Run("localhost:5137")
}

func getVersion(c *gin.Context) {
	game_id := c.Param("game_id")

	config, exists := c.Get("config")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
		return
	}

	version, err := helpers.GetVersion(config.(helpers.Config), game_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"version": version})
}

func getUnits(c *gin.Context) {
	config, exists := c.Get("config")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
		return
	}

	units, err := helpers.GetUnits(config.(helpers.Config))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, units)
}
