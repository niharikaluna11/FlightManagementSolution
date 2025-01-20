package handlers

import (
	"flight-app/config"
	"flight-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoutesBySource(c *gin.Context) {
	var routes []models.Routes
	db, _ := config.ConnectDB()

	sourceId := c.Query("source_id")

	query := db.Preload("Source").Preload("Destination")

	if sourceId != "" {
		query = query.Where("source_id = ?", sourceId)
	}
	if err := query.Find(&routes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch routes"})
		return
	}

	if len(routes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No routes found for the given source_id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"routes": routes})
}

func GetRoutes(ctx *gin.Context) {
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	db = db.Preload("Source").Preload("Destination")

	var routes []models.Routes
	if err := db.Find(&routes).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve routes"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"routes": routes})
}

func CreateRoutes(ctx *gin.Context) {
	var newRoutes models.Routes
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	if err := ctx.ShouldBindJSON(&newRoutes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}
	db = db.Preload("Source").Preload("Destination")
	if err := db.Create(&newRoutes).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save routes"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "routes created successfully",
		"airport": newRoutes,
	})
}
