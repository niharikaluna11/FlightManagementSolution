package handlers

import (
	"flight-app/config"
	"flight-app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAirports(ctx *gin.Context) {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("database error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})

		return
	}

	var airports []models.Airport
	if err := db.Find(&airports).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve airports"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"airports": airports})
}

func GetAirportsWithId(ctx *gin.Context) {
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	var airports []models.Airport
	if err := db.Find(&airports).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve airports"})
		return
	}

	idStr := ctx.DefaultQuery("id", "")

	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID query parameter is required"})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for _, a := range airports {
		if uint(id) == a.ID {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "Airport with ID not found",
		"id":      idStr,
	})

}

func PostAirports(ctx *gin.Context) {
	var newAirport models.Airport
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	if err := ctx.ShouldBindJSON(&newAirport); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	if err := db.Create(&newAirport).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save airport"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Airport created successfully",
		"airport": newAirport,
	})
}
