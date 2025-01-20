package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	ID           uint   `gorm:"not null;unique"`
	FlightNumber string `gorm:"size:20;not null;"` // (e.g., AA101,XY101)
	ModelNumber  string `gorm:"size:255;not null;unique"`
	AirlineName  string `gorm:"size:255;not null;"`
	TotalSeats   uint   `gorm:"not null"`

	ScheduledDeparture time.Time `gorm:"not null"`
	ScheduledArrival   time.Time `gorm:"not null"`
	BasePrice          float64   `gorm:"not null"`
	Status             string    `gorm:"size:50;not null"`

	DepartureAirportID uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ArrivalAirportID   uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	DepartureAirport Airport `gorm:"foreignKey:DepartureAirportID"`
	ArrivalAirport   Airport `gorm:"foreignKey:ArrivalAirportID"`

	DetailClasses []DetailClass `gorm:"foreignKey:FlightID"`
}
