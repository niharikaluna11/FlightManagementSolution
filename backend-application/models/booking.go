package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	FlightID         uint        `gorm:"not null"`
	Flight           Flight      `gorm:"foreignKey:FlightID"`
	DetailsClassID   uint        `gorm:"not null"`
	DetailClass      DetailClass `gorm:"foreignKey:DetailsClassID"`
	BookingReference string      `gorm:"size:50;unique;not null"`
	TotalAmount      float64     `gorm:"not null"`
	Passengers       []Passenger `gorm:"foreignKey:BookingID"`
	Payments         []Payment   `gorm:"foreignKey:BookingID"`
}
