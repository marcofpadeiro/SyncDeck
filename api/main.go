package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/marcofpadeiro/SyncDeck/helpers"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Panic("Error reading config")
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	})

	router.GET("/download/:id", download)
	router.GET("/version/:id", getVersion)
	router.GET("/units", getUnits)
	router.POST("/createUnit", createUnit)

	router.Run("localhost:5137")
}

func getVersion(c *gin.Context) {
	id := c.Param("id")

	config, exists := c.Get("config")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
		return
	}

	version, err := helpers.GetVersion(config.(Config).Save_path+"/metadata.json", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"version": version})
}

func download(c *gin.Context) {
	id := c.Param("id")

	config, exists := c.Get("config")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
		return
	}

	c.File(config.(Config).Save_path + "/" + id + ".zip")
}

func getUnits(c *gin.Context) {
	config, exists := c.Get("config")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
		return
	}

	units, err := helpers.GetUnits(config.(Config).Save_path + "/metadata.json")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, units)
}

func createUnit(c *gin.Context) {
	config, exists := c.Get("config")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Config not found"})
		return
	}

	err := c.Request.ParseMultipartForm(100 << 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse form"})
		return
	}

	// Get the file from the form data
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving file from form data"})
		return
	}
	defer file.Close()

	temp := strings.Split(handler.Filename, ".")
	unit_id := temp[0]
	units, err := helpers.GetUnits(config.(Config).Save_path + "/metadata.json")
	if helpers.CheckExists(units, unit_id) != -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already exists"})
		return
	}

	// Create a new file on the server
	outFile, err := os.Create(filepath.Join(config.(Config).Save_path, handler.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer outFile.Close()

	// Copy the uploaded file to the server
	_, err = io.Copy(outFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error copying file to server"})
		return
	}

	err = helpers.AddUnit(config.(Config).Save_path+"/metadata.json", helpers.Unit{ID: unit_id, Version: 1})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unit added successfully"})
}
