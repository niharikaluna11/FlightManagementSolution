package models

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model
	BookingID      uint    `gorm:"not null"`             
	Booking        Booking `gorm:"foreignKey:BookingID"` 
	Title          Title   `gorm:"type:enum('Mr','Mrs','Miss');not null"`
	FirstName      string  `gorm:"size:100;not null"` 
	LastName       string  `gorm:"size:100;not null"`
	DateOfBirth    string  `gorm:"not null"`         
	SeatNumber     string  `gorm:"size:10;not null"`  
}

type Title string

const (
	Mr   Title = "Mr"
	Mrs  Title = "Mrs"
	Miss Title = "Miss"
)

type PassengerInput struct {
	Title       string `json:"title"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	SeatNumber  string `json:"seat_number"`
}