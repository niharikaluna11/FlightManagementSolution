package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	BookingID     uint    `gorm:"not null;unique"`
	Booking       Booking `gorm:"foreignKey:BookingID"`
	Amount        float64 `gorm:"not null"`
	PaymentMethod string  `gorm:"size:50;not null"`
	TransactionID string  `gorm:"size:100;unique;not null"`
}
