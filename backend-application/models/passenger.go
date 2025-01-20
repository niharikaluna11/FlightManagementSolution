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
