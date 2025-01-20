package models

import "gorm.io/gorm"

type DetailClass struct {
	gorm.Model
	FlightID         uint    `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DetailName       string  `gorm:"size:50;not null"` // Fare class name (e.g., Saver, Flexi)
	Description      string  `gorm:"size:255"`
	BaggageAllowance uint    `gorm:"not null"`
	IsRefundable     bool    `gorm:"not null"`
	PriceMultiplier  float64 `gorm:"not null"`
}
