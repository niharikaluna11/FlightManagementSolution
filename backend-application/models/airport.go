package models

type Airport struct {
	ID      uint   `gorm:"primaryKey"`
	Code    string `gorm:"size:5;unique;not null"`
	Name    string `gorm:"size:255;not null"`
	City    string `gorm:"size:100;not null"`
	State   string `gorm:"size:100;not null"`
	Country string `gorm:"size:100;not null"`
}
