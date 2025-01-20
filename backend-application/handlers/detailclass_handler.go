package handlers

import (
	"flight-app/config"
	"flight-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDetailClass(ctx *gin.Context) {
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	var details []models.DetailClass
	if err := db.Find(&details).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve details"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"details": details})
}

func CreateDetailClass(ctx *gin.Context) {
	var newdetails models.DetailClass
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	if err := ctx.ShouldBindJSON(&newdetails); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	if err := db.Create(&newdetails).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save newdetails"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "newdetails created successfully",
		"airport": newdetails,
	})
}
