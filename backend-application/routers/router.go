package routers

import (
	"flight-app/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Flight application running"})
	})

	//  airport routes
	r.GET("/get-airports", handlers.GetAirports)

	r.GET("/get-airport-by-id", handlers.GetAirportsWithId)

	r.POST("/post-airports", handlers.PostAirports)

	// routes
	r.GET("/get-all-routes", handlers.GetRoutes)
	r.POST("/post-routes", handlers.CreateRoutes)
	r.GET("/get-routes-by-src", handlers.GetRoutesBySource)

	// detail class
	r.GET("/get-detail-class", handlers.GetDetailClass)
	r.POST("/post-detail-class", handlers.CreateDetailClass)

	// flight routes
	r.GET("/get-flights", handlers.GetAllFlights)
	r.POST("/post-flights", handlers.PostFlight)

	r.GET("/search-flights-oneway", handlers.SearchFlightsOneWay)
	r.GET("/search-flights-twoway", handlers.SearchFlightsTwoWay)

	// booking routes
	r.POST("/create-booking", handlers.CreateBooking)
	r.GET("/get-booking", handlers.GetAllBookings)
	r.GET("/get-booking-by-ref", handlers.GetBookingByReference)
	r.GET("/get-bookingpdf", handlers.GetPdf)

	//SendEmailWithPdf
	r.POST("/send-email", handlers.SendEmailWithPdf)

	return r
}
