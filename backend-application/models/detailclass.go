package models

import (
	"errors"

	"gorm.io/gorm"
)

type DetailClass struct {
	gorm.Model
	FlightID         uint    `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DetailName       string  `gorm:"size:50;not null"` // Fare class name (e.g., Saver, Flexi)
	Description      string  `gorm:"size:255"`
	BaggageAllowance uint    `gorm:"not null"`
	IsRefundable     bool    `gorm:"not null"`
	PriceMultiplier  float64 `gorm:"not null"`
	DetailClassID    uint    `gorm:"-"` // Virtual field to handle ID logic, not stored in DB
}

// GORM hook to set DetailClassID before creating a record
func (d *DetailClass) BeforeCreate(tx *gorm.DB) (err error) {
	switch d.DetailName {
	case "Saver":
		d.DetailClassID = 1
	case "Flexi":
		d.DetailClassID = 2
	default:
		return errors.New("invalid DetailName, must be 'Saver' or 'Flexi'")
	}
	return nil
}
