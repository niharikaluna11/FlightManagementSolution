package handlers

import (
	"flight-app/config"
	"flight-app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllFlights(ctx *gin.Context) {
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	db = db.Preload("DepartureAirport").Preload("ArrivalAirport").Preload("DetailClasses")

	var flights []models.Flight
	if err := db.Find(&flights).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve flights"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"flights": flights})
}

func PostFlight(ctx *gin.Context) {
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	var flight models.Flight
	if err := ctx.ShouldBindJSON(&flight); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	var route models.Routes
	if err := db.Where("source_id = ? AND destination_id = ?", flight.DepartureAirportID, flight.ArrivalAirportID).
		First(&route).Error; err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Route does not exist for the given source and destination"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check route existence"})
		}
		return
	}

	db = db.Preload("DepartureAirport").Preload("ArrivalAirport").Preload("DetailClasses")

	if err := db.Create(&flight).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flight"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Flight created successfully", "flight": flight})
}

func GetFlightByID(ctx *gin.Context) {
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	id := ctx.Query("id")
	var flight models.Flight

	db = db.Preload("DepartureAirport").Preload("ArrivalAirport").Preload("DetailClasses")

	if err := db.First(&flight, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"flight": flight})
}

// SearchFlightsTwoWay
func SearchFlightsOneWay(ctx *gin.Context) {
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	src := ctx.Query("source_id")
	dest := ctx.Query("dest_id")
	dateOfJourney := ctx.Query("date_of_journey")

	if src == "" || dest == "" || dateOfJourney == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Source, destination, and date of journey are required"})
		return
	}

	journeyDate, err := time.Parse("2006-01-02", dateOfJourney)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	flights, err := fetchFlights(db, src, dest, journeyDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flights"})
		return
	}

	if len(flights) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"flights": flights})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"flights": flights})
}

func fetchFlights(db *gorm.DB, src, dest string, journeyDate time.Time) ([]models.Flight, error) {
	var flights []models.Flight
	startOfDay := journeyDate
	endOfDay := journeyDate.Add(24 * time.Hour).Add(-time.Second)

	err := db.Preload("DepartureAirport").
		Preload("ArrivalAirport").
		Preload("DetailClasses").
		Where("departure_airport_id = ? AND arrival_airport_id = ? AND scheduled_departure BETWEEN ? AND ?", src, dest, startOfDay, endOfDay).
		Find(&flights).Error

	return flights, err
}

// SearchFlightsTwoWay
func SearchFlightsTwoWay(ctx *gin.Context) {
	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	// Retrieve query parameters
	src := ctx.Query("source_id")
	dest := ctx.Query("dest_id")
	dateOfJourney := ctx.Query("date_of_journey")
	dateOfReturn := ctx.Query("date_of_return")

	// Validate required parameters
	if src == "" || dest == "" || dateOfJourney == "" || dateOfReturn == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Source, destination, date_of_journey, and date_of_return are required"})
		return
	}

	// Parse dates
	journeyDate, err := time.Parse("2006-01-02", dateOfJourney)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for date_of_journey. Use YYYY-MM-DD"})
		return
	}
	returnDate, err := time.Parse("2006-01-02", dateOfReturn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format for date_of_return. Use YYYY-MM-DD"})
		return
	}

	// Fetch flights for the onward journey
	flights1, err := fetchFlights(db, src, dest, journeyDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch onward flights"})
		return
	}

	// Fetch flights for the return journey
	flights2, err := fetchFlights(db, dest, src, returnDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch return flights"})
		return
	}

	// Return flight results
	if len(flights1) == 0 && len(flights2) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No flights available for the selected route and dates"})
		return
	}

	// Combine flights and respond
	result := gin.H{"outbound_flights": flights1, "return_flights": flights2}

	// If flights1 or flights2 are empty, still send response with an appropriate message
	if len(flights1) == 0 {
		result["outbound_flights"] = "No outbound flights available for the selected date"
	}
	if len(flights2) == 0 {
		result["return_flights"] = "No return flights available for the selected date"
	}

	ctx.JSON(http.StatusOK, result)
}
